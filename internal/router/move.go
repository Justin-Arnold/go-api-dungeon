package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
	"github.com/Justin-Arnold/go-api-dungeon/internal/middleware"
	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func HandleMove(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		http.NotFound(w, r)
		return
	}

	moveDirection := dungeon.Direction(pathParts[2])
	switch moveDirection {
	case dungeon.DirectionUp, dungeon.DirectionDown, dungeon.DirectionLeft, dungeon.DirectionRight:
	default:
		http.Error(w, "Invalid direction. Must be up, down, left, or right", http.StatusBadRequest)
		return
	}

	d := dungeon.CreateDungeon()

	state, error := session.GetGameState(r)
	if error != nil {
		http.Error(w, "Game state not found", http.StatusInternalServerError)
		return
	}

	// Check if movement is possible
	if !d.CanMove(state.CurrentRoom, moveDirection) {
		middleware.RedirectToError(w, r, "/error/invalid-direction", http.StatusTemporaryRedirect)
		return
	}
	newRoomID := d.Rooms[state.CurrentRoom].Connections[moveDirection]

	newRoom, exists := d.Rooms[newRoomID]
	if !exists {
		http.Error(w, "Room not found", http.StatusInternalServerError)
		return
	}

	state.CurrentRoom = newRoomID
	state.CurrentRoomType = newRoom.Type

	switch content := newRoom.Content.(type) {
	case dungeon.CombatContent:
		state.CurrentEnemy = content.EnemyType
		state.CurrentEnemyHP = content.HP
		state.CurrentEnemyDamage = content.Damage
		state.CurrentEnemyMaxHP = content.HP
	case dungeon.EventContent:
		state.CurrentEvent = content.EventType
		state.EventDescription = content.Description
		state.EventChoices = content.Choices
	case dungeon.TreasureContent:
		state.TreasureName = content.EventTreasure.Name
	default:
		fmt.Print("BAD")
	}

	//redirect to the new room
	session.SetGameState(w, r, state)

	redirectCookie := &http.Cookie{
		Name:     "redirect-token",
		Value:    "true",
		Path:     "/",
		MaxAge:   10, // Short-lived
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, redirectCookie)

	http.Redirect(w, r, fmt.Sprintf("/room/%s", newRoomID), http.StatusSeeOther)
}

package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
	"github.com/Justin-Arnold/go-api-dungeon/internal/middleware"
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
		// Direction is valid, continue processing
	default:
		http.Error(w, "Invalid direction. Must be up, down, left, or right", http.StatusBadRequest)
		return
	}

	d := dungeon.CreateDungeon()

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")

		gameState, ok := middleware.GetGameState(r.Context())
		if !ok {
			http.Error(w, "Failed to get game state", http.StatusInternalServerError)
			return
		}

		fmt.Print(gameState.CurrentRoom)

		// Check if movement is possible
		if !d.CanMove(gameState.CurrentRoom, moveDirection) {
			fmt.Print("Cannot move")
			middleware.RedirectToError(w, r, "/error/invalid-direction", http.StatusTemporaryRedirect)
			return
		}
		fmt.Print("Can move")
		// Get the new room ID from the current room's connections
		newRoomID := d.Rooms[gameState.CurrentRoom].Connections[moveDirection]

		fmt.Print(d.Rooms[gameState.CurrentRoom])
		// Get the new room
		newRoom, exists := d.Rooms[newRoomID]
		if !exists {
			http.Error(w, "Room not found", http.StatusInternalServerError)
			return
		}

		fmt.Print(newRoom)

		gameState.CurrentRoom = newRoomID
		gameState.CurrentRoomType = newRoom.Type

		fmt.Print(newRoom.Content)

		switch content := newRoom.Content.(type) {
		case dungeon.CombatContent:
			gameState.CurrentEnemy = content.EnemyType
			gameState.CurrentEnemyHP = content.HP
			gameState.CurrentEnemyDamage = content.Damage
			gameState.CurrentEnemyMaxHP = content.HP
		case dungeon.EventContent:
			gameState.CurrentEvent = content.EventType
		case dungeon.TreasureContent:
			gameState.TreasureName = content.EventTreasure.Name
		default:
			fmt.Print("BAD")
		}
		fmt.Print("test")
		if err := json.NewEncoder(w).Encode(gameState); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
		// http.Redirect(w, r, "/room/"+newRoomID, http.StatusTemporaryRedirect)
	}

	data := &TemplateData{
		MoveDirection: moveDirection,
	}

	RenderTemplate(w, "move", data)
}

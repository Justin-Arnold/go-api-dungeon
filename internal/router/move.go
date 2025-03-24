package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
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
	fmt.Print(1)

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Print(2)

		currentRoomID := r.Header.Get("X-Current-Room")
		fmt.Print(9)
		// Check if movement is possible
		if !d.CanMove(currentRoomID, moveDirection) {
			http.Error(w, "Cannot move in that direction", http.StatusBadRequest)
			return
		}

		fmt.Print(3)
		// Get the new room ID from the current room's connections
		newRoomID := d.Rooms[currentRoomID].Connections[moveDirection]
		fmt.Print(4)
		// Get the new room
		newRoom, exists := d.Rooms[newRoomID]
		if !exists {
			http.Error(w, "Room not found", http.StatusInternalServerError)
			return
		}

		// Return game state as JSON
		gameState := map[string]interface{}{
			"currentRoom":     newRoom.ID,
			"currentRoomType": newRoom.Type,
			"redirectTo":      fmt.Sprintf("/room/%s", newRoom.ID),
		}

		switch content := newRoom.Content.(type) {
		case dungeon.CombatContent:
			gameState["currentEnemy"] = content.EnemyType
			gameState["currentEnemyHP"] = content.HP
			gameState["currentEnemyDamage"] = content.Damage
			gameState["currentEnemyMaxHP"] = content.HP
			fmt.Print("Start1")
		case dungeon.EventContent:
			gameState["currentEvent"] = content.EventType
		case dungeon.TreasureContent:
			gameState["treasureName"] = content.EventTreasure.Name
		default:
			fmt.Print("BAD")
		}

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

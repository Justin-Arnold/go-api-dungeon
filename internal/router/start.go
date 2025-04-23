package router

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func HandleStart(w http.ResponseWriter, r *http.Request) {

	gameState, error := session.GetGameState(r)
	if error != nil {
		http.Error(w, "Error retrieving game state", http.StatusInternalServerError)
		return
	}

	gameState.CompletedRooms = append(gameState.CompletedRooms, "start")
	session.SetGameState(w, r, gameState)

	data := &TemplateData{
		CharacterName:  gameState.CharacterName,
		CharacterClass: gameState.CharacterClass,
	}

	RenderTemplate(w, "start", data)
}

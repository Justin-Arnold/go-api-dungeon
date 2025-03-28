package router

import (
	"encoding/json"
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/middleware"
)

func HandleStart(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")

		gameState, ok := middleware.GetGameState(r.Context())
		if !ok {
			// Handle the case where game state is not available
			http.Error(w, "Game state not available", http.StatusBadRequest)
			return
		}

		gameState.CompletedRooms = []string{"start"}

		if err := json.NewEncoder(w).Encode(gameState); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	RenderTemplate(w, "start", &TemplateData{})
}

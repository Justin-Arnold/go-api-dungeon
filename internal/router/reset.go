package router

import (
	"encoding/json"
	"net/http"
)

func HandleReset(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")

		gameState := map[string]interface{}{
			"CharacterName":  nil,
			"CurrentRoom":    nil,
			"CharacterClass": nil,
		}

		if err := json.NewEncoder(w).Encode(gameState); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	RenderTemplate(w, "reset", nil)
}

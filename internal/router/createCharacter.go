package router

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
)

func HandleCreateCharacter(w http.ResponseWriter, r *http.Request) {
	// Extract the character name from the URL path
	// Split the path into parts: ["", "create-character", "name"]
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		http.NotFound(w, r)
		return
	}

	characterName := pathParts[2]
	if characterName == "" {
		http.Error(w, "Character name is required", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")

		// Return game state as JSON
		gameState := map[string]interface{}{
			"CharacterName": characterName,
			"CurrentRoom":   "start",
			// Add any other initial state you want
		}

		if err := json.NewEncoder(w).Encode(gameState); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	data := &TemplateData{
		CharacterName: characterName,
		Classes: []dungeon.ClassInfo{
			{
				Type:        dungeon.ClassShinobi,
				Description: dungeon.GetClassDescription(dungeon.ClassShinobi),
			},
		},
	}

	RenderTemplate(w, "create-character", data)
}

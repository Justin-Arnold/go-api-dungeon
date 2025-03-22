package router

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
)

func HandleChooseClass(w http.ResponseWriter, r *http.Request) {

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		http.NotFound(w, r)
		return
	}

	className := pathParts[2]
	if className == "" {
		http.Error(w, "Class name is required", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")

		classType := dungeon.ClassType(className)

		// get class info
		classStats := dungeon.BaseStats[classType]

		// Return game state as JSON
		gameState := map[string]interface{}{
			"characterClass":  className,
			"characterDamage": classStats.Damage,
			// Add any other initial state you want
		}

		if err := json.NewEncoder(w).Encode(gameState); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	data := &TemplateData{
		CharacterClass: dungeon.ClassType(className),
	}

	// Handle template rendering errors
	RenderTemplate(w, "choose-class", data)
}

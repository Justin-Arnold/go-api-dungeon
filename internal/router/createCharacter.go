package router

import (
	"net/http"
	"strings"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func HandleCreateCharacter(w http.ResponseWriter, r *http.Request) {
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

	gameState, error := session.GetGameState(r)
	if error != nil {
		http.Error(w, "Error retrieving game state", http.StatusInternalServerError)
		return
	}

	gameState.CharacterName = characterName
	gameState.CurrentRoom = "start"

	session.SetGameState(w, r, gameState)

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

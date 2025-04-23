package router

import (
	"net/http"
	"strings"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
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

	classType := dungeon.ClassType(className)
	classStats := dungeon.BaseStats[classType]

	gameState, error := session.GetGameState(r)
	if error != nil {
		http.Error(w, "Error retrieving game state", http.StatusInternalServerError)
		return
	}
	gameState.CharacterClass = className
	gameState.CharacterDamage = classStats.Damage

	session.SetGameState(w, r, gameState)

	data := &TemplateData{
		CharacterClass: dungeon.ClassType(className),
	}

	RenderTemplate(w, "choose-class", data)
}

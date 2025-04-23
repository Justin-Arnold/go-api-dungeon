package router

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func HandleAttack(w http.ResponseWriter, r *http.Request) {

	state, error := session.GetGameState(r)
	if error != nil {
		http.Error(w, "Game state not found", http.StatusInternalServerError)
		return
	}

	newEnemyHP := state.CurrentEnemyHP - state.CharacterDamage

	if newEnemyHP <= 0 {
		newEnemyHP = 0
		state.CompletedRooms = append(state.CompletedRooms, state.CurrentRoom)
	}

	state.CurrentEnemyHP = newEnemyHP

	RenderTemplate(w, "attack", &TemplateData{})
}

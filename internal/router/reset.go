package router

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func HandleReset(w http.ResponseWriter, r *http.Request) {
	session.SetGameState(w, r, session.GameState{})

	RenderTemplate(w, "reset", nil)
}

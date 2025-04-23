package middleware

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func RequireCharacterName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gameState, error := session.GetGameState(r)
		if error != nil {
			http.Error(w, "Error retrieving character name", http.StatusInternalServerError)
			return
		}

		if gameState.CharacterName == "" {
			RedirectToError(w, r, "missing-name", http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}

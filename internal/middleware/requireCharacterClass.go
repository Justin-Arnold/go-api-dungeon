package middleware

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func RequireCharacterClass(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gameState, error := session.GetGameState(r)
		if error != nil {
			http.Error(w, "Error retrieving character class", http.StatusInternalServerError)
			return
		}
		if gameState.CharacterClass == "" {
			// When redirecting to an error page
			RedirectToError(w, r, "missing-class", http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func RequireStart(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		state, error := session.GetGameState(r)
		if error != nil {
			http.Error(w, "Game state not found", http.StatusInternalServerError)
			return
		}

		// TODO use slices.contains instead of a loop
		for _, room := range state.CompletedRooms {
			if room == "start" {
				next.ServeHTTP(w, r)
				return
			}
		}

		RedirectToError(w, r, "not-started", http.StatusTemporaryRedirect)
	})
}

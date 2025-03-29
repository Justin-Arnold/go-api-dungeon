package middleware

import (
	"net/http"
)

func RequireStart(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Accept") == "application/json" {
			gameState, ok := GetGameState(r.Context())
			if !ok {
				// Handle the case where game state is not available
				http.Error(w, "Game state not available", http.StatusBadRequest)
				return
			}

			// check if "start" is in the completed rooms
			for _, room := range gameState.CompletedRooms {
				if room == "start" {
					// If the room is completed, allow the request to proceed
					next.ServeHTTP(w, r)
					return
				}
			}

			RedirectToError(w, r, "not-started", http.StatusTemporaryRedirect)
		}

		next.ServeHTTP(w, r)
	})
}

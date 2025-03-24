package middleware

import (
	"net/http"
)

/*
RequireNoProgression is a middleware that checks if the user has progressed
into the game. If the user has progressed, the middleware will redirect the
user to their current room. If the user has not progressed, the middleware
will pass the request to the next handler.
*/
func RequireNoProgression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Accept") == "application/json" {

			gameState := GameState{
				CurrentRoom: r.Header.Get("X-Current-Room"),
			}

			if gameState.CurrentRoom != "" && gameState.CurrentRoom != "start" {
				RedirectToError(w, r, "invalid-command", http.StatusTemporaryRedirect)
				return
			}
		}

		// For non-API requests, just pass through
		next.ServeHTTP(w, r)
	})
}

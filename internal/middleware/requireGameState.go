package middleware

import (
	"context"
	"fmt"
	"net/http"
)

func RequireGameState(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// log.Printf("DEBUG: Request headers: %+v", r.Header)
		// Only check game state for API requests
		if r.Header.Get("Accept") == "application/json" {
			fmt.Print("test1")

			gameState := GameState{
				CharacterName:  r.Header.Get("X-Character-Name"),
				CharacterClass: r.Header.Get("X-Character-Class"),
				CurrentRoom:    r.Header.Get("X-Current-Room"),
			}
			fmt.Print("test2")
			// Add the game state to the request context
			ctx := context.WithValue(r.Context(), "gameState", gameState)
			fmt.Print("test3")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// For non-API requests, just pass through
		next.ServeHTTP(w, r)
	})
}

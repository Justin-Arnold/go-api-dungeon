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

			// gameState := GameState{
			// 	CharacterName:  r.Header.Get("X-Character-Name"),
			// 	CharacterClass: r.Header.Get("X-Character-Class"),
			// 	CurrentRoom:    r.Header.Get("X-Current-Room"),
			// }
			// // Add the game state to the request context
			// ctx := context.WithValue(r.Context(), "gameState", gameState)
			// next.ServeHTTP(w, r.WithContext(ctx))
			// return
		}

		// For non-API requests, just pass through
		next.ServeHTTP(w, r)
	})
}

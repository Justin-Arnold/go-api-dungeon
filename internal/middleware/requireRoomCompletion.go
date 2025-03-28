package middleware

import (
	"net/http"
)

/*
RequireRoomCompletion is a middleware function that checks if the room is complete.
If the room is not complete, the middleware will block the action and redirect the
user an error page. If the room is complete, the middleware will pass the request
to the next handler.
*/
func RequireRoomCompletion(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Accept") == "application/json" {
			gameState, ok := GetGameState(r.Context())
			if !ok {
				// Handle the case where game state is not available
				http.Error(w, "Game state not available", http.StatusBadRequest)
				return
			}

			//parse the completed rooms, it will be a comma separated string
			//if the current room is in the list of completed rooms, then the room is complete
			for _, room := range gameState.CompletedRooms {
				if room == gameState.CurrentRoom {
					next.ServeHTTP(w, r)
					return
				}
			}

			RedirectToError(w, r, "room-not-complete", http.StatusTemporaryRedirect)
		}

		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

/*
RequireRoomCompletion is a middleware function that checks if the room is complete.
If the room is not complete, the middleware will block the action and redirect the
user an error page. If the room is complete, the middleware will pass the request
to the next handler.
*/
func RequireRoomCompletion(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		state, error := session.GetGameState(r)
		if error != nil {
			http.Error(w, "Game state not found", http.StatusInternalServerError)
			return
		}

		for _, room := range state.CompletedRooms {
			if room == state.CurrentRoom {
				next.ServeHTTP(w, r)
				return
			}
		}

		RedirectToError(w, r, "room-not-complete", http.StatusTemporaryRedirect)
	})
}

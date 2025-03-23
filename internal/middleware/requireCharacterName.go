package middleware

import (
	"fmt"
	"net/http"
)

func RequireCharacterName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Accept") == "application/json" {

			characterName := r.Header.Get("X-Character-Name")

			if characterName == "" {
				// When redirecting to an error page
				fmt.Println("Character name not found")
				RedirectToError(w, r, "missing-name", http.StatusTemporaryRedirect)
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		// For non-API requests, just pass through
		next.ServeHTTP(w, r)
	})
}

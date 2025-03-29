package middleware

import (
	"net/http"
	"strings"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
)

func RequireValidClass(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Accept") == "application/json" {

			pathParts := strings.Split(r.URL.Path, "/")
			if len(pathParts) != 3 {
				http.NotFound(w, r)
				return
			}
			characterClass := pathParts[2]

			//check if the class is valid
			validClass := dungeon.VerifyClass(characterClass)
			if !validClass {
				// When redirecting to an error page
				RedirectToError(w, r, "invalid-class", http.StatusTemporaryRedirect)
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		// For non-API requests, just pass through
		next.ServeHTTP(w, r)
	})
}

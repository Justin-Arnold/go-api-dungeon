package middleware

import (
	"net/http"
)

func RequireCharacterClass(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Accept") == "application/json" {

			characterClass := r.Header.Get("X-Character-Class")

			if characterClass == "" {
				// When redirecting to an error page
				RedirectToError(w, r, "missing-class", http.StatusTemporaryRedirect)
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		// For non-API requests, just pass through
		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RequireCharacterClass(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Accept") == "application/json" {

			gameState, ok := r.Context().Value("gameState").(GameState)
			fmt.Print("test")
			if !ok {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Location", "/start")
				w.WriteHeader(http.StatusTemporaryRedirect) // 307

				json.NewEncoder(w).Encode(map[string]interface{}{
					"redirect": "/start",
					"status":   "redirect_required",
					// "missing": []string{
					// 	// This helps the client understand what's missing
					// 	getMissingFields(hasName, hasClass),
					// },
				})
				return
			}

			if gameState.CharacterClass == "" {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Location", "/start")
				w.WriteHeader(http.StatusTemporaryRedirect) // 307

				json.NewEncoder(w).Encode(map[string]interface{}{
					"redirect": "/start",
					"status":   "redirect_required",
					// "missing": []string{
					// 	// This helps the client understand what's missing
					// 	getMissingFields(hasName, hasClass),
					// },
				})
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		// For non-API requests, just pass through
		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"encoding/json"
	"net/http"
)

func RequireCharacterName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") == "application/json" {
			gameState, ok := r.Context().Value("gameState").(GameState)
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

			// Check if character name exists
			if gameState.CharacterName == "" {

				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Location", "/error")
				w.WriteHeader(http.StatusTemporaryRedirect) // 307

				json.NewEncoder(w).Encode(map[string]interface{}{
					"redirect": "/error",
					//we need to include what middleware caused the error
					"middleware": "RequireCharacterName",
					"status":     "redirect_required",

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

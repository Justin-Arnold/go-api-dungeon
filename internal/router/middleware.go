// internal/router/middleware.go
package router

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type GameState struct {
	CharacterName  string
	CharacterClass string
	CurrentRoom    string
	// Add other state as needed
}

func RequireGameState(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// log.Printf("DEBUG: Request headers: %+v", r.Header)
		// Only check game state for API requests
		if r.Header.Get("Accept") == "application/json" {

			gameState := GameState{
				CharacterName:  r.Header.Get("X-Character-Name"),
				CharacterClass: r.Header.Get("X-Character-Class"),
				CurrentRoom:    r.Header.Get("X-Current-Room"),
			}

			currentPath := r.URL.Path
			splitPath := strings.Split(currentPath, "/")
			rootPath := splitPath[1]
			isChooseClassPath := rootPath == "choose-class"

			hasName := gameState.CharacterName != ""
			hasClass := gameState.CharacterClass != ""
			if !hasName {
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
			}

			if !hasClass && !isChooseClassPath {
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
			}

			// Add the game state to the request context
			ctx := context.WithValue(r.Context(), "gameState", gameState)
			next(w, r.WithContext(ctx))
			return
		}

		// For non-API requests, just pass through
		next(w, r)
	}
}

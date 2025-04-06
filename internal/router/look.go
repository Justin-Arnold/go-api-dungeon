package router

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/middleware"
)

func HandleLook(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")

		gameState, ok := middleware.GetGameState(r.Context())
		if !ok {
			// Handle the case where game state is not available
			http.Error(w, "Game state not available", http.StatusBadRequest)
			return
		}

		//redirect to game state current room
		if gameState.CurrentRoom == "start" {
			redirectCookie := &http.Cookie{
				Name:     "redirect-token",
				Value:    "true",
				Path:     "/",
				MaxAge:   10, // Short-lived
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
			}
			http.SetCookie(w, redirectCookie)
			http.Redirect(w, r, "/start", http.StatusTemporaryRedirect)
			return
		} else {
			redirectCookie := &http.Cookie{
				Name:     "redirect-token",
				Value:    "true",
				Path:     "/",
				MaxAge:   100, // Short-lived
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
			}
			http.SetCookie(w, redirectCookie)
			http.Redirect(w, r, "/room/"+gameState.CurrentRoom, http.StatusTemporaryRedirect)
			return
		}
	}

	RenderTemplate(w, "look", &TemplateData{})
}

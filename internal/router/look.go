package router

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func HandleLook(w http.ResponseWriter, r *http.Request) {

	state, error := session.GetGameState(r)
	if error != nil {
		http.Error(w, "Game state not found", http.StatusInternalServerError)
		return
	}

	if state.CurrentRoom == "start" {
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
			MaxAge:   10, // Short-lived
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, redirectCookie)
		http.Redirect(w, r, "/room/"+state.CurrentRoom, http.StatusTemporaryRedirect)
		return
	}
}

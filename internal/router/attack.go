package router

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func HandleAttack(w http.ResponseWriter, r *http.Request) {

	state, error := session.GetGameState(r)
	if error != nil {
		http.Error(w, "Game state not found", http.StatusInternalServerError)
		return
	}

	newEnemyHP := state.CurrentEnemyHP - state.CharacterDamage

	if newEnemyHP <= 0 {
		newEnemyHP = 0
		state.CompletedRooms = append(state.CompletedRooms, state.CurrentRoom)
	}

	state.CurrentEnemyHP = newEnemyHP

	session.SetGameState(w, r, state)

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
}

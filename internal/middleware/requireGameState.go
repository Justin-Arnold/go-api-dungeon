package middleware

import (
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

// GameState holds the current game state data
type GameState struct {
	CharacterName      string
	CharacterClass     string
	CurrentRoom        string
	CurrentRoomType    dungeon.RoomType
	CurrentEvent       string
	CurrentEnemy       string
	CurrentEnemyDamage int
	TreasureName       string
	CharacterDamage    int
	CurrentEnemyHP     int
	CurrentEnemyMaxHP  int
	CompletedRooms     []string
	EventDescription   string
	EventChoices       []dungeon.Option
}

func RequireGameState(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gameState, error := session.GetGameState(r)
		if error != nil {
			session.SetGameState(w, r, gameState)
		}

		next.ServeHTTP(w, r)
	})
}

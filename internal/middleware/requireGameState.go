package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
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
}

// Define a custom type for context keys
type contextKey string

// GameStateKey is the key used to store the game state in the context
const GameStateKey contextKey = "gameState"

func RequireGameState(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		convertedCharacterDamage, err := strconv.Atoi(r.Header.Get("X-Character-Damage"))
		if err != nil {
			convertedCharacterDamage = 0
		}
		convertedCurrentEnemyHP, err := strconv.Atoi(r.Header.Get("X-Current-Enemy-Hp"))
		if err != nil {
			convertedCurrentEnemyHP = 0
		}

		// log.Printf("DEBUG: Request headers: %+v", r.Header)
		// Only check game state for API requests
		if r.Header.Get("Accept") == "application/json" {
			gameState := GameState{
				CharacterName:   r.Header.Get("X-Character-Name"),
				CharacterClass:  r.Header.Get("X-Character-Class"),
				CurrentRoom:     r.Header.Get("X-Current-Room"),
				CharacterDamage: convertedCharacterDamage,
				CurrentEnemyHP:  convertedCurrentEnemyHP,
				CompletedRooms:  strings.Split(r.Header.Get("X-Completed-Rooms"), ","),
			}
			// Add the game state to the request context
			ctx := context.WithValue(r.Context(), GameStateKey, gameState)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// For non-API requests, just pass through
		next.ServeHTTP(w, r)
	})
}

func GetGameState(ctx context.Context) (GameState, bool) {
	gameState, ok := ctx.Value(GameStateKey).(GameState)
	return gameState, ok
}

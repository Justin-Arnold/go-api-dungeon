package session

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
)

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

// Define a custom type for context keys
type contextKey string

// GameStateKey is the key used to store the game state in the context
const GameStateKey contextKey = "gameState"

/*
GetGameState retrieves the GameState struct from the session.
It returns a pointer to the GameState and an error if session retrieval fails.
If no GameState exists in the session, it returns a new, zeroed GameState.
*/
func GetGameState(r *http.Request) (GameState, error) {
	session, ok := Get[GameState](r, string(GameStateKey))
	if !ok {
		// Log underlying file system or decryption errors
		log.Printf("GetGameState: Failed to retrieve session")
		// Return a new empty state AND the error, allowing caller to decide how critical the error is
		// Alternatively, just return nil, err depending on desired behavior
		return GameState{}, fmt.Errorf("failed to get session")
	}

	// Success! Return the existing GameState pointer
	return session, nil
}

/*
SetGameState sets the GameState struct in the session.
It returns an error if session storage fails.
If the session is successfully stored, it returns nil.
*/
func SetGameState(w http.ResponseWriter, r *http.Request, state GameState) error {
	err := Set(w, r, string(GameStateKey), state)

	if err != nil {
		return err
	}
	return nil
}

func ResetGameState(w http.ResponseWriter, r *http.Request) error {
	// Reset the game state to a new instance
	newState := GameState{}

	// Store the new game state in the session
	err := SetGameState(w, r, newState)
	if err != nil {
		return err
	}

	return nil
}

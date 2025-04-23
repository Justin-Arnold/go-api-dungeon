package session

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
	"github.com/gorilla/sessions"
)

// Session name constant
const SessionName = "game-session"

// WARNING: Use secure random keys in production!
// Generate random keys (recommended for production)
// authKey := securecookie.GenerateRandomKey(64)
// encryptionKey := securecookie.GenerateRandomKey(32)

// For demonstration purposes ONLY - DO NOT use these static keys in production
var authKey = []byte("your-very-secret-authentication-key-64-bytes-long-or-more") // Min 32 bytes ideally 64
var encryptionKey = []byte("this-is-exactly-32-bytes-long!!!")                    // 16, 24, or 32 bytes for AES

// Define the directory to store session files
const sessionDir = "./sessions" // Make sure this directory exists and is writable!

// Create the FilesystemStore
// Pass the directory, auth key, and encryption key
var store = sessions.NewFilesystemStore(sessionDir, authKey, encryptionKey)

func init() {
	gob.Register(dungeon.ClassType(""))
	gob.Register(GameState{})
	// Create the directory if it doesn't exist
	err := os.MkdirAll(sessionDir, 0700) // 0700: Owner can read/write/execute
	if err != nil {
		log.Fatalf("Failed to create session directory: %v", err)
	}

	store.Options = &sessions.Options{
		Path:     "/",                  // Cookie available for all paths
		MaxAge:   86400 * 7,            // Session lasts for 7 days (in seconds)
		HttpOnly: true,                 // Prevent JavaScript access to the cookie
		Secure:   false,                // Set to true if using HTTPS (RECOMMENDED in production)
		SameSite: http.SameSiteLaxMode, // Good default for security (Lax or Strict)
	}
}

// Get retrieves a value from the session
func Get[T any](r *http.Request, key string) (T, bool) {
	var zero T

	session, err := store.Get(r, SessionName) // Renamed 'state' to 'session' for clarity
	if err != nil {
		// Log the error when session retrieval fails
		log.Printf("session.Get: Failed to retrieve session: %v", err)
		return zero, false // Indicate key not found (because session failed)
	}

	// Look up the value in the session map
	value, ok := session.Values[key]
	if !ok {
		// Key does not exist in the session map
		log.Printf("session.Get: Key '%s' not found in session", key) // Optional: Log missing key
		return zero, false
	}

	typedValue, typeOK := value.(T)
	if !typeOK {
		// Key exists, but the value stored is NOT of the expected type T.
		// This often indicates a programming error (saving as one type, retrieving as another).
		log.Printf("session.Get: Type mismatch for key '%s'. Expected %T but got %T", key, zero, value)
		// The type assertion returns the zero value of T when it fails (`typedValue` will be zero T here),
		// so we return typedValue and false.
		return zero, false // Return zero T and false on type mismatch
	}

	// The key was found, and the value was successfully asserted to type T.
	return typedValue, true
}

// Set stores a value in the session
func Set(w http.ResponseWriter, r *http.Request, key string, value interface{}) error {
	session, err := store.Get(r, SessionName)
	if err != nil {
		return err
	}

	session.Values[key] = value

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

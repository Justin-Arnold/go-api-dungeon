package middleware

import (
	"net/http"
)

type GameState struct {
	CharacterName  string
	CharacterClass string
	CurrentRoom    string
}

type Middleware func(next http.Handler) http.Handler

func Register(h http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	var handler http.Handler = h

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	return handler.ServeHTTP

}

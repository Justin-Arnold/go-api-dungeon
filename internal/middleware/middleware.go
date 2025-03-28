package middleware

import (
	"net/http"
	"strings"
)

type Middleware func(next http.Handler) http.Handler

func Register(h http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	var handler http.Handler = h

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	return handler.ServeHTTP

}

// RedirectToError sets the redirect cookie and redirects to an error page
// This ensures that error pages accessed via redirect are allowed by the BlockDirectAccess middleware
func RedirectToError(w http.ResponseWriter, r *http.Request, errorPath string, status int) {
	// Set the redirect cookie
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

	// Ensure the path starts with /error/
	if !strings.HasPrefix(errorPath, "/error/") {
		errorPath = "/error/" + strings.TrimPrefix(errorPath, "/")
	}

	http.Redirect(w, r, errorPath, http.StatusTemporaryRedirect)
}

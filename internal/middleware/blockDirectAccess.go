package middleware

import (
	"net/http"
	"strings"
)

/*
BlockDirectAccess is a middleware that blocks direct access to a given
endpoint. These routes can be redirected to, but not accessed directly.
This is for pages like an error page, where the user should only be able
to access it through a redirect but not by navigating to the URL directly.
*/
func BlockDirectAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if this is an error route
		if strings.HasPrefix(r.URL.Path, "/error/") {
			// Check for our redirect cookie
			cookie, err := r.Cookie("redirect-token")

			if err == nil && cookie.Value == "true" {
				// This is a legitimate redirect, clear the cookie and allow access
				clearCookie := &http.Cookie{
					Name:     "redirect-token",
					Value:    "",
					Path:     "/",
					MaxAge:   -1,
					HttpOnly: true,
					Secure:   true,
					SameSite: http.SameSiteStrictMode,
				}
				http.SetCookie(w, clearCookie)

				// Allow access to the error page
				next.ServeHTTP(w, r)
				return
			}

			// This appears to be direct access, set a cookie and redirect to the no-direct-access error
			redirectCookie := &http.Cookie{
				Name:     "redirect-token",
				Value:    "true",
				Path:     "/",
				MaxAge:   10, // Very short-lived cookie
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
			}
			http.SetCookie(w, redirectCookie)

			// Redirect to a specific error page for no direct access
			http.Redirect(w, r, "/error/invalid-command", http.StatusTemporaryRedirect)
			return
		}

		// For non-error routes, set the redirect cookie before redirecting to an error page
		if r.Header.Get("X-Error-Redirect") == "true" {
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
		}

		// For all other requests, just pass through
		next.ServeHTTP(w, r)
	})
}

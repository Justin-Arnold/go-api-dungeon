package middleware

import (
	"fmt"
	"net/http"
)

/*
BlockDirectAccess is a middleware that blocks direct access to a given
endpoint. These routes can be redirected to, but not accessed directly.
This is for pages like an error page, where the user should only be able
to access it through a redirect but not by navigating to the URL directly.
*/
func BlockDirectAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Accept") == "application/json" {
			next.ServeHTTP(w, r)
			return
		}

		// Check if this is an error route
		// if strings.HasPrefix(r.URL.Path, "/error/") {
		// Check for our redirect cookie
		cookie, err := r.Cookie("redirect-token")

		fmt.Println("Cookie value:", cookie)
		fmt.Println("Error:", err)

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

			next.ServeHTTP(w, r)
			return
		}

		// // If this is the "invalid-command" page and we're already coming from a redirect,
		// // break the potential loop and just show the page
		// if strings.Contains(r.URL.Path, "/error/invalid-command") &&
		// 	strings.Contains(r.Referer(), "/error/") {
		// 	next.ServeHTTP(w, r)
		// 	return
		// }

		// Otherwise redirect to the no-direct-access error
		RedirectToError(w, r, "/error/invalid-command", http.StatusTemporaryRedirect)
		// }
	})
}

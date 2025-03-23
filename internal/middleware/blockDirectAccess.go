package middleware

import (
	"fmt"
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

			fmt.Println("Error route accessed", cookie)
			fmt.Println("Error route accessed", err)

			if err == nil && cookie.Value == "true" {
				fmt.Println("Redirect token found", cookie)
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

				fmt.Println("Redirect token cleared", w)

				// Allow access to the error page
				fmt.Println("Error route accessed directly", cookie)
				next.ServeHTTP(w, r)
				return
			}

			// // If this is the "invalid-command" page and we're already coming from a redirect,
			// // break the potential loop and just show the page
			// if strings.Contains(r.URL.Path, "/error/invalid-command") &&
			// 	strings.Contains(r.Referer(), "/error/") {
			// 	fmt.Println("Preventing redirect loop, showing invalid-command page")
			// 	next.ServeHTTP(w, r)
			// 	return
			// }

			// Otherwise redirect to the no-direct-access error
			RedirectToError(w, r, "/error/invalid-command", http.StatusTemporaryRedirect)
			return
		}

		// For all other requests, just pass through
		next.ServeHTTP(w, r)
	})
}

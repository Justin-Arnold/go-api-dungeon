package router

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleError(w http.ResponseWriter, r *http.Request) {

	// get page we are coming from
	referer := r.Header.Get("Referer")

	if referer == "" {
		referer = "/"
	}

	fmt.Println("Referrer: ", referer)

	pathParts := strings.Split(r.URL.Path, "/")

	fmt.Println("Path parts: ", pathParts)
	fmt.Println("Path parts length: ", len(pathParts))
	if len(pathParts) < 3 {
		http.NotFound(w, r)
		return
	}

	fmt.Println("Path parts: ", pathParts)

	data := &TemplateData{}

	fmt.Println("Data: ", data)

	errorType := pathParts[2]

	fmt.Println("Error type: ", errorType)
	fmt.Println(errorType == "missing-class")

	switch errorType {
	case "invalid-command":
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
		RenderTemplate(w, "errors/invalid-command", data)
		return
	case "missing-class":
		fmt.Println("Missing class!")
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
		RenderTemplate(w, "errors/missing-class", data)
		return
	case "missing-name":
		fmt.Println("Missing name!")
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
		RenderTemplate(w, "errors/missing-name", data)
		return
	default:
		// RenderTemplate(w, "errors/404", data)
		return
	}
}

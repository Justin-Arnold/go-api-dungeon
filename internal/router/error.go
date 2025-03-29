package router

import (
	"net/http"
	"strings"
)

func HandleError(w http.ResponseWriter, r *http.Request) {

	// get page we are coming from
	referer := r.Header.Get("Referer")

	if referer == "" {
		referer = "/"
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 3 {
		http.NotFound(w, r)
		return
	}

	data := &TemplateData{}

	errorType := pathParts[2]

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
	case "room-not-complete":
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
		RenderTemplate(w, "errors/room-not-complete", data)
		return
	case "missing-class":
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
	case "not-started":
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
		RenderTemplate(w, "errors/not-started", data)
		return
	default:
		// RenderTemplate(w, "errors/404", data)
		return
	}
}

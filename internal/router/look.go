package router

import "net/http"

func HandleLook(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "look", &TemplateData{})
}

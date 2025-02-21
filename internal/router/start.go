package router

import "net/http"

func HandleStart(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "start", &TemplateData{})
}

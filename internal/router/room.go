package router

import (
	"net/http"
)

func HandleRoom(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")

		// currentRoom := r.Header.Get("X-Current-Room")

		// data := &TemplateData{
		// 	CurrentEnemy: currentEnemy,
		// }

		// RenderTemplate(w, "room", data)
	}

	data := &TemplateData{}

	RenderTemplate(w, "room", data)
}

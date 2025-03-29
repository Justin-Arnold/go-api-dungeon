package router

import (
	"encoding/json"
	"net/http"
)

func HandleReset(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")

		data := map[string]string{}

		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	RenderTemplate(w, "reset", nil)
}

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

	var errorTitle string

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		http.NotFound(w, r)
		return
	}

	data := &TemplateData{
		ErrorTitle: errorTitle,
		//set whole r as error message, stringified
		ErrorMessage: fmt.Sprintf("%+v", r),
	}

	errorType := pathParts[2]

	switch errorType {
	case "invalid-command":
		RenderTemplate(w, "errors/invalid-command", data)
		return
	default:
		RenderTemplate(w, "errors/404", data)
		return
	}
}

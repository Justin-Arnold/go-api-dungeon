package router

import (
	"fmt"
	"net/http"
)

func HandleError(w http.ResponseWriter, r *http.Request) {

	// get page we are coming from
	referer := r.Header.Get("Referer")

	if referer == "" {
		referer = "/"
	}

	var errorTitle string
	// var errorMessage string

	//switch case for what referer ends in
	// switch {
	// Regex for /choose-class/???
	// case regexp.MustCompile(`/choose-class/\w+$`).MatchString(referer):
	// 	errorTitle = "Choose Class Error"
	// 	errorMessage = strings.Split(referer, "/choose-class/")[1]
	// //default case
	// default:
	// 	errorTitle = "Error"
	// 	errorMessage = "An error occurred"
	// }

	data := &TemplateData{
		ErrorTitle: errorTitle,
		//set whole r as error message, stringified
		ErrorMessage: fmt.Sprintf("%+v", r),
	}

	RenderTemplate(w, "error", data)
}

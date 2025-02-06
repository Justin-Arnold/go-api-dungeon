package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// templateData is a struct to hold any data we want to pass to our templates
type templateData struct {
	// Add fields as needed
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *templateData) {
	// Parse both the layout and the page template
	templates := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/"+tmpl+".html",
	))

	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Check if the path is exactly "/" since this handler will match all paths by default
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	renderTemplate(w, "index", &templateData{})
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "start", &templateData{})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Register route handlers
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/start", handleStart)

	log.Printf("starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

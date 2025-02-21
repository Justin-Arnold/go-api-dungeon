package router

import (
	"html/template"
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
)

type TemplateData struct {
	CharacterName  string
	Character      dungeon.Character
	Classes        []dungeon.ClassInfo
	CharacterClass dungeon.ClassType
	MoveDirection  dungeon.Direction
	CurrentEnemy   string
	RedirectTo     string
	// Add fields as needed
}

func Init() {
	http.HandleFunc("/create-character/", HandleCreateCharacter)
	http.HandleFunc("/choose-class/", RequireGameState(HandleChooseClass))
	http.HandleFunc("/move/", RequireGameState(HandleMove))
	http.HandleFunc("/room/", RequireGameState(HandleRoom))
	http.HandleFunc("/reset/", HandleReset)
	http.HandleFunc("/start", HandleStart)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *TemplateData) {
	// Parse both the layout and the page template
	templates := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/room-types/combat-room.html",
		"templates/room-types/treasure-room.html",
		"templates/room-types/empty-room.html",
		"templates/room-types/event-room.html",
		"templates/"+tmpl+".html",
	))

	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

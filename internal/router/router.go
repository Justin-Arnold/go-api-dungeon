package router

import (
	"html/template"
	"net/http"

	"github.com/Justin-Arnold/go-api-dungeon/internal/dungeon"
	"github.com/Justin-Arnold/go-api-dungeon/internal/middleware"
)

type TemplateData struct {
	CharacterName  string
	Character      dungeon.Character
	Classes        []dungeon.ClassInfo
	CharacterClass dungeon.ClassType
	MoveDirection  dungeon.Direction
	CurrentEnemy   string
	RedirectTo     string
	ErrorTitle     string
	ErrorMessage   string
	// Add fields as needed
}

func Init() {
	http.HandleFunc("/create-character/", HandleCreateCharacter)
	http.HandleFunc("/choose-class/", middleware.Register(HandleChooseClass,
		middleware.RequireGameState,
		middleware.RequireCharacterName,
	))
	http.HandleFunc("/move/", middleware.Register(HandleMove,
		middleware.RequireGameState,
		middleware.RequireCharacterName,
		middleware.RequireCharacterClass,
	))
	http.HandleFunc("/room/", middleware.Register(HandleRoom,
		middleware.RequireGameState,
		middleware.RequireCharacterName,
		middleware.RequireCharacterClass,
	))
	http.HandleFunc("/reset/", HandleReset)
	http.HandleFunc("/start", HandleStart)
	http.HandleFunc("/error", HandleError)
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

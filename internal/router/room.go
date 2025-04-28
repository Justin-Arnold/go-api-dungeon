package router

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/Justin-Arnold/go-api-dungeon/internal/session"
)

func HandleRoom(w http.ResponseWriter, r *http.Request) {

	state, errror := session.GetGameState(r)
	if errror != nil {
		http.Error(w, "Game state not found", http.StatusInternalServerError)
		return
	}

	data := &TemplateData{
		RoomType:               string(state.CurrentRoomType),
		CurrentEnemy:           state.CurrentEnemy,
		CurrentEnemyHP:         state.CurrentEnemyHP,
		CurrentEnemyMaxHP:      state.CurrentEnemyMaxHP,
		EventType:              state.CurrentEvent,
		EventText:              state.EventDescription,
		EventChoices:           getEventChoices(state),
		EnemyImageURL:          getEnemyImageURL(state),
		EnemyHealthInlineStyle: template.CSS(getHealthBarInlineStyle(state)),
	}

	RenderTemplate(w, "room", data)
}

func getHealthBarInlineStyle(state session.GameState) string {
	healthPercentage := (state.CurrentEnemyHP * 100) / state.CurrentEnemyMaxHP
	inlineStyle := "width: " + strconv.Itoa(healthPercentage) + "%;"

	return inlineStyle
}

func getEnemyImageURL(state session.GameState) string {
	enemyType := state.CurrentEnemy
	enemyType = strings.ToLower(enemyType)
	enemyType = strings.ReplaceAll(enemyType, " ", "-")
	concatenatedURL := "/static/enemy-" + enemyType + ".webp"

	return concatenatedURL

}

func getEventChoices(state session.GameState) []string {
	choices := make([]string, len(state.EventChoices))
	for i, choice := range state.EventChoices {
		choices[i] = choice.Text
	}
	return choices
}

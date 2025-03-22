package router

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func HandleAttack(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")

		attackDamage := r.Header.Get("X-Character-Damage")
		currentEnemyHP := r.Header.Get("X-Current-Enemy-Hp")

		parsedAttackDamage, _ := strconv.Atoi(attackDamage)
		parsedCurrentEnemyHP, _ := strconv.Atoi(currentEnemyHP)

		newEnemyHP := parsedCurrentEnemyHP - parsedAttackDamage

		if newEnemyHP < 0 {
			newEnemyHP = 0
		}

		gameState := map[string]interface{}{
			"currentEnemyHP": newEnemyHP,
		}

		if err := json.NewEncoder(w).Encode(gameState); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return

	}

	RenderTemplate(w, "attack", &TemplateData{})
}

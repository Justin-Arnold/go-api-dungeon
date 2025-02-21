package dungeon

import (
	"math/rand"
)

type Treasure struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Quality     TreasureQuality `json:"tier"`
	Effect      StatEffect      `json:"Effect"`
}

type TreasureQuality string

const (
	Profane TreasureQuality = "profane"
)

func GetRandomTreasureByQuality(quality TreasureQuality) *Treasure {
	var allTreasure []Treasure = []Treasure{
		{
			Name:        "Sacrificial Blade of the Cult",
			Description: "",
			Quality:     Profane,
			Effect: StatEffect{
				Name:        "Profane Strength",
				Description: "",
				Stat:        "Damage",
			},
		},
	}

	var validTreasure []Treasure
	for _, treasure := range allTreasure {
		if treasure.Quality == quality {
			validTreasure = append(validTreasure, treasure)
		}
	}

	return &validTreasure[rand.Intn(len(validTreasure))]
}

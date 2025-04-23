package dungeon

import "fmt"

// RoomType represents the different types of rooms possible in the dungeon
type RoomType string

const (
	RoomTypeCombat   RoomType = "COMBAT"
	RoomTypeEmpty    RoomType = "EMPTY"
	RoomTypeEvent    RoomType = "EVENT"
	RoomTypeTreasure RoomType = "TREASURE"
)

// Direction represents possible movement directions between rooms
type Direction string

const (
	DirectionUp    Direction = "up"
	DirectionDown  Direction = "down"
	DirectionLeft  Direction = "left"
	DirectionRight Direction = "right"
)

// Room represents a single room in the dungeon
type Room struct {
	ID          string               `json:"id"`
	Type        RoomType             `json:"type"`
	Connections map[Direction]string `json:"connections"` // maps direction to room ID
	Content     interface{}          `json:"content"`     // varies based on room type
}

type CombatContent struct {
	EnemyType   string `json:"enemyType"`
	Description string `json:"description"`
	HP          int    `json:"hp"`
	Damage      int    `json:"damage"`
	Rewards     Reward `json:"rewards"`
}

type StatEffect struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stat        string `json:"stat"`
	Change      int    `json:"change"`
	ChangeType  string `json:"changeType"`
}

type EventContent struct {
	EventType   string   `json:"eventType"`
	Description string   `json:"description"`
	Choices     []Option `json:"choices"`
}

type TreasureContent struct {
	EventType     string   `json:"eventType"`
	Description   string   `json:"description"`
	EventTreasure Treasure `json:"treasure"`
}

// Option represents a choice in an event room
type Option struct {
	Text   string `json:"text"`
	Reward Reward `json:"reward"`
}

// Reward represents possible rewards from combat or events
type Reward struct {
	Gold        int          `json:"gold"`
	Experience  int          `json:"experience"`
	StatEffects []StatEffect `json:"statEffects"`
}

// Dungeon represents the full dungeon structure
type Dungeon struct {
	Rooms map[string]*Room `json:"rooms"`
}

// NewDungeon creates a new dungeon instance
func NewDungeon() *Dungeon {
	return &Dungeon{
		Rooms: make(map[string]*Room),
	}
}

// CanMove checks if movement in a direction is possible
func (d *Dungeon) CanMove(currentRoomID string, dir Direction) bool {
	fmt.Print("CanMove")
	room, exists := d.Rooms[currentRoomID]
	if !exists {
		return false
	}

	_, hasConnection := room.Connections[dir]
	return hasConnection
}

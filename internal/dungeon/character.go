package dungeon

import (
	"errors"
	"math"
)

// ClassType represents the available character classes
type ClassType string

type ClassInfo struct {
	Type        ClassType
	Description string
}

const (
	ClassShinobi ClassType = "Shinobi"
)

// Stats represents a character's base statistics
type Stats struct {
	CurrentHP   int     `json:"currentHP"`
	MaxHP       int     `json:"maxHp"`
	Damage      int     `json:"strength"`
	Defense     int     `json:"intelligence"`
	Speed       int     `json:"dexterity"`
	DodgeChance float64 `json:"dodgeChance"`
	CritChance  float64 `json:"critChance"`
	Luck        float64 `json:"luck"`
}

type Character struct {
	Name  string    `json:"name"`
	Class ClassType `json:"class"`
	Level int       `json:"level"`
	Stats Stats     `json:"stats"`
	Gold  int       `json:"gold"`
	XP    int       `json:"xp"`
}

var BaseStats = map[ClassType]Stats{
	ClassShinobi: {
		CurrentHP:   90,   // Base 100
		MaxHP:       90,   // Base 100
		Damage:      11,   // Base 10
		Defense:     9,    // Base 10
		Speed:       15,   // Base 10
		DodgeChance: 0.15, // Base 0,10
		CritChance:  0.20, // Base 0.10
		Luck:        1.00, //base 1.00
	},
}

func NewCharacter(name string, class ClassType) (*Character, error) {
	stats, ok := BaseStats[class]
	if !ok {
		return nil, errors.New("invalid character class")
	}

	return &Character{
		Name:  name,
		Class: class,
		Level: 1,
		Stats: stats,
		Gold:  0,
		XP:    0,
	}, nil
}

func (c *Character) XPToNextLevel() int {
	return c.Level * 100
}

func (c *Character) CanLevelUp() bool {
	return c.XP >= c.XPToNextLevel()
}

func (c *Character) LevelUp() error {
	if !c.CanLevelUp() {
		return errors.New("not enough XP to level up")
	}

	c.Level++
	c.XP -= c.XPToNextLevel()

	// Increase stats based on class
	switch c.Class {
	case ClassShinobi:
		c.Stats.MaxHP += int(math.Floor(float64(c.Stats.MaxHP) * ((.1 * c.Stats.Luck) * float64(c.Level))))
	}

	// Heal to full on level up
	c.Stats.CurrentHP = c.Stats.MaxHP
	return nil
}

// TakeDamage applies damage to the character
func (c *Character) TakeDamage(amount int) {
	c.Stats.CurrentHP -= amount
	if c.Stats.CurrentHP < 0 {
		c.Stats.CurrentHP = 0
	}
}

// Heal recovers HP up to max
func (c *Character) Heal(amount int) {
	c.Stats.CurrentHP += amount
	if c.Stats.CurrentHP > c.Stats.MaxHP {
		c.Stats.CurrentHP = c.Stats.MaxHP
	}
}

// IsAlive checks if character is still alive
func (c *Character) IsAlive() bool {
	return c.Stats.CurrentHP > 0
}

// GetClassDescription returns a description of the class
func GetClassDescription(class ClassType) string {
	descriptions := map[ClassType]string{
		ClassShinobi: "Shinobi are fast, agile fighters with high dodge and crit chances. They excel at dealing consistent damage while avoiding hits.",
	}
	return descriptions[class]
}

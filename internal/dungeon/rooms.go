package dungeon

// Example of creating the dungeon structure
func CreateDungeon() *Dungeon {
	d := NewDungeon()

	// Start room
	d.Rooms["start"] = &Room{
		ID:   "start",
		Type: RoomTypeEmpty,
		Connections: map[Direction]string{
			DirectionRight: "room1",
		},
	}

	// Combat Room
	d.Rooms["room1"] = &Room{
		ID:   "room1",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionLeft:  "start",
			DirectionRight: "room2",
		},
		Content: CombatContent{
			EnemyType:   "goblin",
			Description: "A fight",
			HP:          20,
			Damage:      5,
			Rewards:     Reward{},
		},
	}

	d.Rooms["room2"] = &Room{
		ID:   "room2",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionLeft: "room1",
			DirectionUp:   "room3",
			DirectionDown: "room4",
		},
		Content: CombatContent{
			EnemyType: "large goblin",
			HP:        30,
			Damage:    7,
			Rewards:   Reward{},
		},
	}

	d.Rooms["room3"] = &Room{
		ID:   "room3",
		Type: RoomTypeEvent,
		Connections: map[Direction]string{
			DirectionDown:  "room2",
			DirectionRight: "room5",
		},
		Content: EventContent{
			EventType:   "Choice",
			Description: "Filler",
		},
	}

	d.Rooms["room4"] = &Room{
		ID:   "room4",
		Type: RoomTypeEvent,
		Connections: map[Direction]string{
			DirectionUp:    "room2",
			DirectionRight: "room6",
		},
		Content: EventContent{
			EventType:   "Choice",
			Description: "Filler",
		},
	}

	d.Rooms["room5"] = &Room{
		ID:   "room5",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionUp:    "room7",
			DirectionLeft:  "room3",
			DirectionRight: "room9",
		},
		Content: CombatContent{
			EnemyType: "Goblin Shaman",
			HP:        10,
			Damage:    15,
			Rewards:   Reward{},
		},
	}

	d.Rooms["room6"] = &Room{
		ID:   "room6",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionDown:  "room8",
			DirectionLeft:  "room4",
			DirectionRight: "room10",
		},
		Content: CombatContent{
			EnemyType: "Goblin Warrior",
			HP:        25,
			Damage:    12,
			Rewards:   Reward{},
		},
	}

	d.Rooms["room7"] = &Room{
		ID:   "room7",
		Type: RoomTypeTreasure,
		Connections: map[Direction]string{
			DirectionDown: "room5",
		},
		Content: TreasureContent{
			EventType:   "Legendary",
			Description: "A legendary reward",
		},
	}

	d.Rooms["room8"] = &Room{
		ID:   "room8",
		Type: RoomTypeTreasure,
		Connections: map[Direction]string{
			DirectionUp: "room6",
		},
		Content: TreasureContent{
			EventType:   "Legendary",
			Description: "A legendary reward",
		},
	}

	d.Rooms["room9"] = &Room{
		ID:   "room9",
		Type: RoomTypeEvent,
		Connections: map[Direction]string{
			DirectionDown:  "room11",
			DirectionLeft:  "room5",
			DirectionRight: "room13",
		},
		Content: EventContent{
			EventType:   "Chaos",
			Description: "Filler",
		},
	}

	d.Rooms["room10"] = &Room{
		ID:   "room10",
		Type: RoomTypeEvent,
		Connections: map[Direction]string{
			DirectionUp:    "room11",
			DirectionLeft:  "room6",
			DirectionRight: "room14",
		},
		Content: EventContent{
			EventType:   "Chaos",
			Description: "Filler",
		},
	}

	d.Rooms["room11"] = &Room{
		ID:   "room11",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionUp:    "room9",
			DirectionDown:  "room10",
			DirectionRight: "room12",
		},
		Content: CombatContent{
			EnemyType: "Chaos Amalgam",
			HP:        40,
			Damage:    15,
			Rewards:   Reward{},
		},
	}

	d.Rooms["room12"] = &Room{
		ID:   "room12",
		Type: RoomTypeTreasure,
		Connections: map[Direction]string{
			DirectionLeft: "room11",
		},
		Content: TreasureContent{
			EventType:   "Chaos",
			Description: "A truly incredible reward",
		},
	}

	d.Rooms["room13"] = &Room{
		ID:   "room13",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionLeft:  "room9",
			DirectionRight: "room15",
		},
		Content: CombatContent{
			EnemyType: "Forest Spirit",
			HP:        15,
			Damage:    15,
			Rewards:   Reward{},
		},
	}

	d.Rooms["room14"] = &Room{
		ID:   "room14",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionLeft:  "room10",
			DirectionRight: "room16",
		},
		Content: CombatContent{
			EnemyType: "Forest Druid",
			HP:        30,
			Damage:    10,
			Rewards:   Reward{},
		},
	}

	d.Rooms["room15"] = &Room{
		ID:   "room15",
		Type: RoomTypeEvent,
		Connections: map[Direction]string{
			DirectionLeft:  "room13",
			DirectionRight: "room17",
		},
		Content: EventContent{
			EventType:   "fate",
			Description: "this or that",
		},
	}

	d.Rooms["room16"] = &Room{
		ID:   "room16",
		Type: RoomTypeEvent,
		Connections: map[Direction]string{
			DirectionLeft:  "room14",
			DirectionRight: "room19",
		},
		Content: EventContent{
			EventType:   "fate",
			Description: "this or that",
		},
	}

	d.Rooms["room17"] = &Room{
		ID:   "room17",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionLeft: "room15",
			DirectionDown: "room18",
		},
		Content: CombatContent{
			EnemyType: "Cult Advocate",
			HP:        25,
			Damage:    25,
			Rewards:   Reward{},
		},
	}

	d.Rooms["room18"] = &Room{
		ID:   "room18",
		Type: RoomTypeEvent,
		Connections: map[Direction]string{
			DirectionUp:    "room17",
			DirectionRight: "room20",
			DirectionDown:  "room19",
		},
		Content: EventContent{
			EventType:   "Chaos",
			Description: "Filler",
		},
	}

	d.Rooms["room19"] = &Room{
		ID:   "room19",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionLeft: "room16",
			DirectionUp:   "room18",
		},
		Content: CombatContent{
			EnemyType: "Cult Fanatic",
			HP:        20,
			Damage:    30,
			Rewards:   Reward{},
		},
	}

	d.Rooms["room20"] = &Room{
		ID:   "room20",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionLeft:  "room18",
			DirectionRight: "room22",
			DirectionDown:  "room21",
			DirectionUp:    "boss",
		},
		Content: CombatContent{
			EnemyType: "Cult Leader",
			HP:        40,
			Damage:    20,
			Rewards:   Reward{},
		},
	}

	d.Rooms["room21"] = &Room{
		ID:   "room21",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionUp: "room20",
		},
		Content: CombatContent{
			EnemyType: "Cult Congregation",
			HP:        60,
			Damage:    10,
			Rewards:   Reward{},
		},
	}

	d.Rooms["room22"] = &Room{
		ID:   "room22",
		Type: RoomTypeTreasure,
		Connections: map[Direction]string{
			DirectionLeft: "room20",
		},
		Content: TreasureContent{
			EventType:     "Profane",
			Description:   "Filler",
			EventTreasure: *GetRandomTreasureByQuality("profane"),
		},
	}

	d.Rooms["boss"] = &Room{
		ID:   "boss",
		Type: RoomTypeCombat,
		Connections: map[Direction]string{
			DirectionDown: "room20",
		},
		Content: CombatContent{
			EnemyType: "Summoned Hell King",
			HP:        100,
			Damage:    30,
			Rewards:   Reward{},
		},
	}

	return d
}

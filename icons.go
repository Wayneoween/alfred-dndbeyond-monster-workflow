package main

import (
	"strings"

	aw "github.com/deanishe/awgo"
)

var (
	// UpdateAvailable is an icon
	UpdateAvailable = &aw.Icon{Value: "icons/update-available.png"}
	// MonsterIconDefault is an icon for monsters of unknown type
	MonsterIconDefault = &aw.Icon{Value: "icons/dnd/default.png"}
	// MonsterIconAbberations is an icon
	MonsterIconAbberations = &aw.Icon{Value: "icons/dnd/aberration.jpg"}
	// MonsterIconBeasts is an icon
	MonsterIconBeasts = &aw.Icon{Value: "icons/dnd/beast.jpg"}
	// MonsterIconCelestials is an icon
	MonsterIconCelestials = &aw.Icon{Value: "icons/dnd/celestial.jpg"}
	// MonsterIconConstructs is an icon
	MonsterIconConstructs = &aw.Icon{Value: "icons/dnd/construct.jpg"}
	// MonsterIconDragons is an icon
	MonsterIconDragons = &aw.Icon{Value: "icons/dnd/dragon.jpg"}
	// MonsterIconElementals is an icon
	MonsterIconElementals = &aw.Icon{Value: "icons/dnd/elemental.jpg"}
	// MonsterIconFey is an icon
	MonsterIconFey = &aw.Icon{Value: "icons/dnd/fey.jpg"}
	// MonsterIconFiends is an icon
	MonsterIconFiends = &aw.Icon{Value: "icons/dnd/fiend.jpg"}
	// MonsterIconGiants is an icon
	MonsterIconGiants = &aw.Icon{Value: "icons/dnd/giant.jpg"}
	// MonsterIconHumanoids is an icon
	MonsterIconHumanoids = &aw.Icon{Value: "icons/dnd/humanoid.jpg"}
	// MonsterIconMonstrosities is an icon
	MonsterIconMonstrosities = &aw.Icon{Value: "icons/dnd/monstrosity.jpg"}
	// MonsterIconOozes is an icon
	MonsterIconOozes = &aw.Icon{Value: "icons/dnd/ooze.jpg"}
	// MonsterIconPlants is an icon
	MonsterIconPlants = &aw.Icon{Value: "icons/dnd/plant.jpg"}
	// MonsterIconUndead is an icon
	MonsterIconUndead = &aw.Icon{Value: "icons/dnd/undead.jpg"}
)

func getIconForType(monsterType string) *aw.Icon {

	var temp *aw.Icon

	switch strings.ToLower(monsterType) {
	case "aberration":
		temp = MonsterIconAbberations
	case "beast":
		temp = MonsterIconBeasts
	case "celestial":
		temp = MonsterIconCelestials
	case "construct":
		temp = MonsterIconConstructs
	case "dragon":
		temp = MonsterIconDragons
	case "elemental":
		temp = MonsterIconElementals
	case "fey":
		temp = MonsterIconFey
	case "fiend":
		temp = MonsterIconFiends
	case "giant":
		temp = MonsterIconGiants
	case "humanoid":
		temp = MonsterIconHumanoids
	case "monstrosity":
		temp = MonsterIconMonstrosities
	case "ooze":
		temp = MonsterIconOozes
	case "plant":
		temp = MonsterIconPlants
	case "undead":
		temp = MonsterIconUndead
	default:
		temp = MonsterIconDefault
	}

	return temp
}

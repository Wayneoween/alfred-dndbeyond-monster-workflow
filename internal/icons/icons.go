// Package icons provides icon management for D&D monster types.
package icons

import (
	"strings"

	aw "github.com/deanishe/awgo"
)

var (
	// UpdateAvailable is the icon shown when a workflow update is available
	UpdateAvailable = &aw.Icon{Value: "icons/update-available.png"}
	// Default is the icon for monsters of unknown type
	Default = &aw.Icon{Value: "icons/dnd/default.png"}
	// Aberrations is the icon for aberration type monsters
	Aberrations = &aw.Icon{Value: "icons/dnd/aberration.jpg"}
	// Beasts is the icon for beast type monsters
	Beasts = &aw.Icon{Value: "icons/dnd/beast.jpg"}
	// Celestials is the icon for celestial type monsters
	Celestials = &aw.Icon{Value: "icons/dnd/celestial.jpg"}
	// Constructs is the icon for construct type monsters
	Constructs = &aw.Icon{Value: "icons/dnd/construct.jpg"}
	// Dragons is the icon for dragon type monsters
	Dragons = &aw.Icon{Value: "icons/dnd/dragon.jpg"}
	// Elementals is the icon for elemental type monsters
	Elementals = &aw.Icon{Value: "icons/dnd/elemental.jpg"}
	// Fey is the icon for fey type monsters
	Fey = &aw.Icon{Value: "icons/dnd/fey.jpg"}
	// Fiends is the icon for fiend type monsters
	Fiends = &aw.Icon{Value: "icons/dnd/fiend.jpg"}
	// Giants is the icon for giant type monsters
	Giants = &aw.Icon{Value: "icons/dnd/giant.jpg"}
	// Humanoids is the icon for humanoid type monsters
	Humanoids = &aw.Icon{Value: "icons/dnd/humanoid.jpg"}
	// Monstrosities is the icon for monstrosity type monsters
	Monstrosities = &aw.Icon{Value: "icons/dnd/monstrosity.jpg"}
	// Oozes is the icon for ooze type monsters
	Oozes = &aw.Icon{Value: "icons/dnd/ooze.jpg"}
	// Plants is the icon for plant type monsters
	Plants = &aw.Icon{Value: "icons/dnd/plant.jpg"}
	// Undead is the icon for undead type monsters
	Undead = &aw.Icon{Value: "icons/dnd/undead.jpg"}
)

// ForType returns the appropriate icon for a given monster type.
// The type comparison is case-insensitive.
// Returns Default icon if the type is not recognized.
func ForType(monsterType string) *aw.Icon {
	switch strings.ToLower(monsterType) {
	case "aberration":
		return Aberrations
	case "beast":
		return Beasts
	case "celestial":
		return Celestials
	case "construct":
		return Constructs
	case "dragon":
		return Dragons
	case "elemental":
		return Elementals
	case "fey":
		return Fey
	case "fiend":
		return Fiends
	case "giant":
		return Giants
	case "humanoid":
		return Humanoids
	case "monstrosity":
		return Monstrosities
	case "ooze":
		return Oozes
	case "plant":
		return Plants
	case "undead":
		return Undead
	default:
		return Default
	}
}

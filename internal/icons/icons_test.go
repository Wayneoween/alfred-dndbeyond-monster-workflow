package icons

import (
	"testing"

	aw "github.com/deanishe/awgo"
)

func TestForType(t *testing.T) {
	tests := []struct {
		name         string
		monsterType  string
		expectedIcon *aw.Icon
	}{
		{
			name:         "aberration lowercase",
			monsterType:  "aberration",
			expectedIcon: Aberrations,
		},
		{
			name:         "aberration uppercase",
			monsterType:  "ABERRATION",
			expectedIcon: Aberrations,
		},
		{
			name:         "aberration mixed case",
			monsterType:  "ABeRrAtIoN",
			expectedIcon: Aberrations,
		},
		{
			name:         "beast",
			monsterType:  "beast",
			expectedIcon: Beasts,
		},
		{
			name:         "celestial",
			monsterType:  "celestial",
			expectedIcon: Celestials,
		},
		{
			name:         "construct",
			monsterType:  "construct",
			expectedIcon: Constructs,
		},
		{
			name:         "dragon",
			monsterType:  "dragon",
			expectedIcon: Dragons,
		},
		{
			name:         "elemental",
			monsterType:  "elemental",
			expectedIcon: Elementals,
		},
		{
			name:         "fey",
			monsterType:  "fey",
			expectedIcon: Fey,
		},
		{
			name:         "fiend",
			monsterType:  "fiend",
			expectedIcon: Fiends,
		},
		{
			name:         "giant",
			monsterType:  "giant",
			expectedIcon: Giants,
		},
		{
			name:         "humanoid",
			monsterType:  "humanoid",
			expectedIcon: Humanoids,
		},
		{
			name:         "monstrosity",
			monsterType:  "monstrosity",
			expectedIcon: Monstrosities,
		},
		{
			name:         "ooze",
			monsterType:  "ooze",
			expectedIcon: Oozes,
		},
		{
			name:         "plant",
			monsterType:  "plant",
			expectedIcon: Plants,
		},
		{
			name:         "undead",
			monsterType:  "undead",
			expectedIcon: Undead,
		},
		{
			name:         "unknown type returns default",
			monsterType:  "unknown",
			expectedIcon: Default,
		},
		{
			name:         "empty string returns default",
			monsterType:  "",
			expectedIcon: Default,
		},
		{
			name:         "swarm returns default",
			monsterType:  "swarm",
			expectedIcon: Default,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ForType(tt.monsterType)
			if result != tt.expectedIcon {
				t.Errorf("ForType(%q) = %v; want %v", tt.monsterType, result, tt.expectedIcon)
			}
		})
	}
}

func TestIconsExist(t *testing.T) {
	icons := map[string]*aw.Icon{
		"UpdateAvailable": UpdateAvailable,
		"Default":         Default,
		"Aberrations":     Aberrations,
		"Beasts":          Beasts,
		"Celestials":      Celestials,
		"Constructs":      Constructs,
		"Dragons":         Dragons,
		"Elementals":      Elementals,
		"Fey":             Fey,
		"Fiends":          Fiends,
		"Giants":          Giants,
		"Humanoids":       Humanoids,
		"Monstrosities":   Monstrosities,
		"Oozes":           Oozes,
		"Plants":          Plants,
		"Undead":          Undead,
	}

	for name, icon := range icons {
		t.Run(name, func(t *testing.T) {
			if icon == nil {
				t.Errorf("%s is nil", name)
			}
			if icon.Value == "" {
				t.Errorf("%s has empty Value", name)
			}
		})
	}
}

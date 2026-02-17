package main

import (
	"testing"

	aw "github.com/deanishe/awgo"
)

func TestGetIconForType(t *testing.T) {
	tests := []struct {
		name         string
		monsterType  string
		expectedIcon *aw.Icon
	}{
		{
			name:         "aberration lowercase",
			monsterType:  "aberration",
			expectedIcon: MonsterIconAbberations,
		},
		{
			name:         "aberration uppercase",
			monsterType:  "ABERRATION",
			expectedIcon: MonsterIconAbberations,
		},
		{
			name:         "aberration mixed case",
			monsterType:  "ABeRrAtIoN",
			expectedIcon: MonsterIconAbberations,
		},
		{
			name:         "beast",
			monsterType:  "beast",
			expectedIcon: MonsterIconBeasts,
		},
		{
			name:         "celestial",
			monsterType:  "celestial",
			expectedIcon: MonsterIconCelestials,
		},
		{
			name:         "construct",
			monsterType:  "construct",
			expectedIcon: MonsterIconConstructs,
		},
		{
			name:         "dragon",
			monsterType:  "dragon",
			expectedIcon: MonsterIconDragons,
		},
		{
			name:         "elemental",
			monsterType:  "elemental",
			expectedIcon: MonsterIconElementals,
		},
		{
			name:         "fey",
			monsterType:  "fey",
			expectedIcon: MonsterIconFey,
		},
		{
			name:         "fiend",
			monsterType:  "fiend",
			expectedIcon: MonsterIconFiends,
		},
		{
			name:         "giant",
			monsterType:  "giant",
			expectedIcon: MonsterIconGiants,
		},
		{
			name:         "humanoid",
			monsterType:  "humanoid",
			expectedIcon: MonsterIconHumanoids,
		},
		{
			name:         "monstrosity",
			monsterType:  "monstrosity",
			expectedIcon: MonsterIconMonstrosities,
		},
		{
			name:         "ooze",
			monsterType:  "ooze",
			expectedIcon: MonsterIconOozes,
		},
		{
			name:         "plant",
			monsterType:  "plant",
			expectedIcon: MonsterIconPlants,
		},
		{
			name:         "undead",
			monsterType:  "undead",
			expectedIcon: MonsterIconUndead,
		},
		{
			name:         "unknown type returns default",
			monsterType:  "unknown",
			expectedIcon: MonsterIconDefault,
		},
		{
			name:         "empty string returns default",
			monsterType:  "",
			expectedIcon: MonsterIconDefault,
		},
		{
			name:         "swarm returns default",
			monsterType:  "swarm",
			expectedIcon: MonsterIconDefault,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getIconForType(tt.monsterType)
			if result != tt.expectedIcon {
				t.Errorf("getIconForType(%q) = %v; want %v", tt.monsterType, result, tt.expectedIcon)
			}
		})
	}
}

func TestIconsExist(t *testing.T) {
	icons := map[string]*aw.Icon{
		"UpdateAvailable":        UpdateAvailable,
		"MonsterIconDefault":     MonsterIconDefault,
		"MonsterIconAbberations": MonsterIconAbberations,
		"MonsterIconBeasts":      MonsterIconBeasts,
		"MonsterIconCelestials":  MonsterIconCelestials,
		"MonsterIconConstructs":  MonsterIconConstructs,
		"MonsterIconDragons":     MonsterIconDragons,
		"MonsterIconElementals":  MonsterIconElementals,
		"MonsterIconFey":         MonsterIconFey,
		"MonsterIconFiends":      MonsterIconFiends,
		"MonsterIconGiants":      MonsterIconGiants,
		"MonsterIconHumanoids":   MonsterIconHumanoids,
		"MonsterIconMonstrosities": MonsterIconMonstrosities,
		"MonsterIconOozes":       MonsterIconOozes,
		"MonsterIconPlants":      MonsterIconPlants,
		"MonsterIconUndead":      MonsterIconUndead,
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

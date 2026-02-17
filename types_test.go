package main

import (
	"encoding/json"
	"testing"
)

func TestMonsterJSONMarshaling(t *testing.T) {
	monster := Monster{
		NameDE:        "Goblin",
		NameDEUlisses: "Goblin Ulisses",
		NameEN:        "Goblin",
		PageDE:        "42",
		PageEN:        "166",
		Src:           []string{"MM", "SRD"},
		SrdName:       "goblin",
		Size:          "Small",
		Type:          "humanoid",
		Tags:          "goblinoid",
		Alignment:     "neutral evil",
		Cr:            "1/4",
		Xp:            "50",
		SingleLine:    "Small humanoid (goblinoid), neutral evil",
	}

	// Test marshaling
	data, err := json.Marshal(monster)
	if err != nil {
		t.Fatalf("Failed to marshal monster: %v", err)
	}

	// Test unmarshaling
	var unmarshaled Monster
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal monster: %v", err)
	}

	// Verify fields
	if unmarshaled.NameEN != monster.NameEN {
		t.Errorf("NameEN = %q; want %q", unmarshaled.NameEN, monster.NameEN)
	}
	if unmarshaled.NameDE != monster.NameDE {
		t.Errorf("NameDE = %q; want %q", unmarshaled.NameDE, monster.NameDE)
	}
	if unmarshaled.Size != monster.Size {
		t.Errorf("Size = %q; want %q", unmarshaled.Size, monster.Size)
	}
	if unmarshaled.Type != monster.Type {
		t.Errorf("Type = %q; want %q", unmarshaled.Type, monster.Type)
	}
	if unmarshaled.Cr != monster.Cr {
		t.Errorf("Cr = %q; want %q", unmarshaled.Cr, monster.Cr)
	}
	if len(unmarshaled.Src) != len(monster.Src) {
		t.Errorf("Src length = %d; want %d", len(unmarshaled.Src), len(monster.Src))
	}
}

func TestD3ResultSetJSONMarshaling(t *testing.T) {
	resultSet := D3ResultSet{
		O: "monster",
		V: 0.7,
		Monster: []Monster{
			{
				NameEN: "Goblin",
				Size:   "Small",
				Type:   "humanoid",
				Cr:     "1/4",
			},
			{
				NameEN: "Orc",
				Size:   "Medium",
				Type:   "humanoid",
				Cr:     "1/2",
			},
		},
	}

	// Test marshaling
	data, err := json.Marshal(resultSet)
	if err != nil {
		t.Fatalf("Failed to marshal resultSet: %v", err)
	}

	// Test unmarshaling
	var unmarshaled D3ResultSet
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal resultSet: %v", err)
	}

	// Verify fields
	if unmarshaled.O != resultSet.O {
		t.Errorf("O = %q; want %q", unmarshaled.O, resultSet.O)
	}
	if unmarshaled.V != resultSet.V {
		t.Errorf("V = %f; want %f", unmarshaled.V, resultSet.V)
	}
	if len(unmarshaled.Monster) != len(resultSet.Monster) {
		t.Errorf("Monster length = %d; want %d", len(unmarshaled.Monster), len(resultSet.Monster))
	}
	if len(unmarshaled.Monster) > 0 {
		if unmarshaled.Monster[0].NameEN != resultSet.Monster[0].NameEN {
			t.Errorf("Monster[0].NameEN = %q; want %q", unmarshaled.Monster[0].NameEN, resultSet.Monster[0].NameEN)
		}
	}
}

func TestMonsterJSONTags(t *testing.T) {
	// Test that JSON tags are correctly defined
	jsonData := `{
		"name_de": "Test DE",
		"name_de_ulisses": "Test Ulisses",
		"name_en": "Test EN",
		"page_de": "10",
		"page_en": "20",
		"src": ["MM"],
		"srdname": "test",
		"size": "Medium",
		"type": "beast",
		"tags": "test-tag",
		"alignment": "neutral",
		"cr": "1",
		"xp": "200",
		"singleline": "Medium beast, neutral"
	}`

	var monster Monster
	err := json.Unmarshal([]byte(jsonData), &monster)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify all fields were properly unmarshaled
	if monster.NameDE != "Test DE" {
		t.Errorf("NameDE = %q; want %q", monster.NameDE, "Test DE")
	}
	if monster.NameDEUlisses != "Test Ulisses" {
		t.Errorf("NameDEUlisses = %q; want %q", monster.NameDEUlisses, "Test Ulisses")
	}
	if monster.NameEN != "Test EN" {
		t.Errorf("NameEN = %q; want %q", monster.NameEN, "Test EN")
	}
	if monster.PageDE != "10" {
		t.Errorf("PageDE = %q; want %q", monster.PageDE, "10")
	}
	if monster.PageEN != "20" {
		t.Errorf("PageEN = %q; want %q", monster.PageEN, "20")
	}
	if len(monster.Src) != 1 || monster.Src[0] != "MM" {
		t.Errorf("Src = %v; want [MM]", monster.Src)
	}
	if monster.SrdName != "test" {
		t.Errorf("SrdName = %q; want %q", monster.SrdName, "test")
	}
	if monster.Size != "Medium" {
		t.Errorf("Size = %q; want %q", monster.Size, "Medium")
	}
	if monster.Type != "beast" {
		t.Errorf("Type = %q; want %q", monster.Type, "beast")
	}
	if monster.Tags != "test-tag" {
		t.Errorf("Tags = %q; want %q", monster.Tags, "test-tag")
	}
	if monster.Alignment != "neutral" {
		t.Errorf("Alignment = %q; want %q", monster.Alignment, "neutral")
	}
	if monster.Cr != "1" {
		t.Errorf("Cr = %q; want %q", monster.Cr, "1")
	}
	if monster.Xp != "200" {
		t.Errorf("Xp = %q; want %q", monster.Xp, "200")
	}
	if monster.SingleLine != "Medium beast, neutral" {
		t.Errorf("SingleLine = %q; want %q", monster.SingleLine, "Medium beast, neutral")
	}
}

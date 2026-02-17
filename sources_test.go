package main

import (
	"testing"
)

func TestExcludeSrcList(t *testing.T) {
	expectedExcluded := []string{
		"AiME-BRF",
		"AiME-Eria",
		"AiME-RIV",
		"AiME-RRF",
		"AiME-SLH",
		"AiME-WdD",
		"AiME-Wild",
		"AVENT-M",
		"AVENT-W",
		"CC",
		"CTH-GHOUL",
		"CTHULHU",
		"D3",
		"MARGREVE",
		"MTGAFR",
		"Myth-AdDM",
		"Myth-Held",
		"Myth-Saga",
		"RAGNAROK",
		"STRANGE",
		"ToB",
		"ToB2",
	}

	if len(excludeSrc) != len(expectedExcluded) {
		t.Errorf("excludeSrc length = %d; want %d", len(excludeSrc), len(expectedExcluded))
	}

	// Create a map for quick lookup
	excludeMap := make(map[string]bool)
	for _, src := range excludeSrc {
		excludeMap[src] = true
	}

	// Check all expected sources are in the exclude list
	for _, expected := range expectedExcluded {
		if !excludeMap[expected] {
			t.Errorf("Expected source %q not found in excludeSrc", expected)
		}
	}
}

func TestExcludeSrcNotEmpty(t *testing.T) {
	if len(excludeSrc) == 0 {
		t.Error("excludeSrc should not be empty")
	}
}

package sources

import "testing"

func TestExcludedSourcesList(t *testing.T) {
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

	if len(ExcludedSources) != len(expectedExcluded) {
		t.Errorf("ExcludedSources length = %d; want %d", len(ExcludedSources), len(expectedExcluded))
	}

	// Create a map for quick lookup
	excludeMap := make(map[string]bool)
	for _, src := range ExcludedSources {
		excludeMap[src] = true
	}

	// Check all expected sources are in the exclude list
	for _, expected := range expectedExcluded {
		if !excludeMap[expected] {
			t.Errorf("Expected source %q not found in ExcludedSources", expected)
		}
	}
}

func TestExcludedSourcesNotEmpty(t *testing.T) {
	if len(ExcludedSources) == 0 {
		t.Error("ExcludedSources should not be empty")
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name   string
		listA  []string
		listB  []string
		expect bool
	}{
		{
			name:   "both lists empty",
			listA:  []string{},
			listB:  []string{},
			expect: false,
		},
		{
			name:   "first list empty",
			listA:  []string{},
			listB:  []string{"a", "b"},
			expect: false,
		},
		{
			name:   "second list empty",
			listA:  []string{"a", "b"},
			listB:  []string{},
			expect: false,
		},
		{
			name:   "no common elements",
			listA:  []string{"a", "b", "c"},
			listB:  []string{"d", "e", "f"},
			expect: false,
		},
		{
			name:   "one common element at start",
			listA:  []string{"a", "b", "c"},
			listB:  []string{"a", "d", "e"},
			expect: true,
		},
		{
			name:   "one common element in middle",
			listA:  []string{"a", "b", "c"},
			listB:  []string{"d", "b", "e"},
			expect: true,
		},
		{
			name:   "one common element at end",
			listA:  []string{"a", "b", "c"},
			listB:  []string{"d", "e", "c"},
			expect: true,
		},
		{
			name:   "multiple common elements",
			listA:  []string{"a", "b", "c"},
			listB:  []string{"a", "b", "d"},
			expect: true,
		},
		{
			name:   "all elements common",
			listA:  []string{"a", "b", "c"},
			listB:  []string{"a", "b", "c"},
			expect: true,
		},
		{
			name:   "real world example with excluded sources",
			listA:  []string{"AiME-BRF", "CTHULHU", "ToB"},
			listB:  []string{"MM", "CTHULHU"},
			expect: true,
		},
		{
			name:   "real world example with no excluded sources",
			listA:  []string{"AiME-BRF", "CTHULHU", "ToB"},
			listB:  []string{"MM", "SRD"},
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsAny(tt.listA, tt.listB)
			if result != tt.expect {
				t.Errorf("ContainsAny(%v, %v) = %v; want %v", tt.listA, tt.listB, result, tt.expect)
			}
		})
	}
}

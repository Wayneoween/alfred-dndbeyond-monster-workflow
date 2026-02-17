package main

import "testing"

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
			result := containsAny(tt.listA, tt.listB)
			if result != tt.expect {
				t.Errorf("containsAny(%v, %v) = %v; want %v", tt.listA, tt.listB, result, tt.expect)
			}
		})
	}
}

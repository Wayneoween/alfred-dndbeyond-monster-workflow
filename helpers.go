package main

// containsAny checks if a string is present in a slice
func containsAny(listA []string, listB []string) bool {
	for _, a := range listA {
		for _, b := range listB {
			if a == b {
				return true
			}
		}
	}
	return false
}

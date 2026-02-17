// Package sources provides source book filtering for D&D monsters.
package sources

// ExcludedSources lists source books that should be excluded from results.
// These are typically non-official or third-party source books.
var ExcludedSources = []string{
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

// ContainsAny checks if any string from listA is present in listB
func ContainsAny(listA []string, listB []string) bool {
	for _, a := range listA {
		for _, b := range listB {
			if a == b {
				return true
			}
		}
	}
	return false
}

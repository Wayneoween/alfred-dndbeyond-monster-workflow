// Package monster provides data types and models for D&D monster information.
package monster

// ResultSet mirrors the response from the dnddeutsch.de API
type ResultSet struct {
	O       string    `json:"o"`
	V       float64   `json:"v"`
	Monster []Monster `json:"monster"`
}

// Monster represents a D&D monster with all its attributes
type Monster struct {
	NameDE        string   `json:"name_de"`
	NameDEUlisses string   `json:"name_de_ulisses"`
	NameEN        string   `json:"name_en"`
	PageDE        string   `json:"page_de"`
	PageEN        string   `json:"page_en"`
	Src           []string `json:"src"`
	SrdName       string   `json:"srdname"`
	Size          string   `json:"size"`
	Type          string   `json:"type"`
	Tags          string   `json:"tags"`
	Alignment     string   `json:"alignment"`
	Cr            string   `json:"cr"`
	Xp            string   `json:"xp"`
	SingleLine    string   `json:"singleline"`
}

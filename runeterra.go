// Package runeterra is an API for Riot Games' Legends of Runeterra API.
package runeterra

// Client for the Riot API.
type Client struct {
	Config Config
}

// Config for the Client.
type Config struct {
	APIKey string
}

// DataDragon contains LoR's game assets and data.
type DataDragon struct {
	Keywords []Keyword `json:"keywords"`
}

// Keyword represents a game play mechanic.
type Keyword struct {
	NameRef     string `json:"nameRef"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Regions are locales of Runeterra, which represent deck archetypes.
type Regions struct {
	NameRef          string `json:"nameRef"`
	Name             string `json:"name"`
	Abbereviation    string `json:"abbreviation"`
	IconAbsolutePath string `json:"iconAbsolutePath"`
}

// SpellSpeeds are the priorities that spells can take.
type SpellSpeeds struct {
	Name    string `json:"name"`
	NameRef string `json:"nameRef"`
}

// Rarities of cards.
type Rarities struct {
	Name    string `json:"name"`
	NameRef string `json:"nameRef"`
}

type Card struct {
	Name string `json:"name"`
}

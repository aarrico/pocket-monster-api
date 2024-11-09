package seed

import (
	"encoding/json"
	"github.com/aarrico/pocket-monster-api/internal/utils"
	"log"
)

type ApiData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ApiResp struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []ApiData `json:"results"`
}

type SubField struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Description struct {
	Description string   `json:"description"`
	Language    SubField `json:"language"`
}

type BasicInfo struct {
	Name         string        `json:"name"`
	Descriptions []Description `json:"descriptions"`
}

type BaseStats struct {
	Value int      `json:"base_stat"`
	Stat  SubField `json:"stat"`
}

type Types struct {
	Slot int32    `json:"slot"`
	Type SubField `json:"type"`
}

type Pokemon struct {
	Name           string      `json:"name"`
	Height         int32       `json:"height"`
	Weight         int32       `json:"weight"`
	Species        SubField    `json:"species"`
	BaseExperience int32       `json:"base_experience"`
	IsDefault      bool        `json:"is_default"`
	SortOrder      int32       `json:"order"`
	BaseStats      []BaseStats `json:"stats"`
	Types          []Types     `json:"types"`
}

type EffectEntry struct {
	Effect   string   `json:"effect"`
	Language SubField `json:"language"`
}

type PokemonForAbility struct {
	IsHidden bool     `json:"is_hidden"`
	Slot     int      `json:"slot"`
	Pokemon  SubField `json:"pokemon"`
}

type Ability struct {
	Name    string              `json:"name"`
	Entries []EffectEntry       `json:"effect_entries"`
	Pokemon []PokemonForAbility `json:"pokemon"`
}
type MoveMeta struct {
	Ailment       SubField `json:"ailment"`
	Category      SubField `json:"category"`
	MinHits       int32    `json:"min_hits"`
	MaxHits       int32    `json:"max_hits"`
	MinTurns      int32    `json:"min_turns"`
	MaxTurns      int32    `json:"max_turns"`
	Drain         int32    `json:"drain"`
	Healing       int32    `json:"healing"`
	CritRate      int32    `json:"crit_rate"`
	AilmentChance int32    `json:"ailment_chance"`
	FlinchChance  int32    `json:"flinch_chance"`
	StatChance    int32    `json:"stat_chance"`
}

type Move struct {
	Name        string        `json:"name"`
	Accuracy    int32         `json:"accuracy"`
	PowerPoints int32         `json:"pp"`
	Priority    int32         `json:"priority"`
	Power       int32         `json:"power"`
	DamageClass SubField      `json:"damage_class"`
	Effect      []EffectEntry `json:"effect_entries"`
	Meta        MoveMeta      `json:"meta"`
	Target      SubField      `json:"target"`
	Type        SubField      `json:"type"`
	Pokemon     []SubField    `json:"learned_by_pokemon"`
}

func populateTableFromApi(url string, populate populateTableFunc) {
	for {
		body := utils.GetBodyFromUrl(url, true)

		var apiResponse ApiResp
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			log.Println("error unmarshalling api response:", err)
			break
		}

		for _, rawData := range apiResponse.Results {
			populate(rawData.Url)
		}

		url = apiResponse.Next
		if url == "" {
			break
		}
	}
}

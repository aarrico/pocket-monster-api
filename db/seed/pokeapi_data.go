package seed

import (
	"encoding/json"
	"fmt"
	"github.com/aarrico/pocket-monster-api/internal/utils"
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

func populateTableFromApi(url string, populate populateTableFunc) {
	for {
		body := utils.GetBodyFromUrl(url, true)

		var apiResponse ApiResp
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			fmt.Println("error unmarshalling api response:", err)
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

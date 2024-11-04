package seed

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/aarrico/pocket-monster-api/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
	url2 "net/url"
	"strconv"
	"strings"
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

type Species struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemon struct {
	Name           string  `json:"name"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Species        Species `json:"species"`
	BaseExperience int     `json:"base_experience"`
	IsDefault      bool    `json:"is_default"`
	SortOrder      int     `json:"order"`
}

func getPokemon(url string) (db.CreatePokemonParams, error) {
	body := utils.GetBodyFromUrl(url, true)

	var pkmn Pokemon
	var pkmnParams db.CreatePokemonParams
	if err := json.Unmarshal(body, &pkmn); err != nil {
		fmt.Println("error unmarshalling pokemon data:", err)
		return pkmnParams, err
	}

	u, _ := url2.Parse(pkmn.Species.Url)
	parts := strings.Split(u.Path, "/")
	dexOrderStr := parts[len(parts)-2]
	dexOrder, _ := strconv.Atoi(dexOrderStr)

	pkmnParams.Name = pgtype.Text{String: pkmn.Name, Valid: true}
	pkmnParams.Height = pgtype.Int4{Int32: int32(pkmn.Height), Valid: true}
	pkmnParams.Weight = pgtype.Int4{Int32: int32(pkmn.Weight), Valid: true}
	pkmnParams.NationalDexOrder = pgtype.Int4{Int32: int32(dexOrder), Valid: true}
	pkmnParams.BaseExperience = pgtype.Int4{Int32: int32(pkmn.BaseExperience), Valid: true}
	pkmnParams.SortOrder = pgtype.Int4{Int32: int32(pkmn.SortOrder), Valid: true}
	pkmnParams.IsDefault = pgtype.Bool{Bool: pkmn.IsDefault, Valid: true}

	return pkmnParams, nil
}

func SeedPokemon() {
	ctx := context.Background()
	pool := utils.ConnectToDb(ctx)
	queries := db.New(pool)

	url := "https://pokeapi.co/api/v2/pokemon"

	for {
		body := utils.GetBodyFromUrl(url, true)

		var apiResponse ApiResp
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			fmt.Println("error unmarshalling api response:", err)
			break
		}

		for _, rawData := range apiResponse.Results {
			pkmn, err := getPokemon(rawData.Url)
			if err != nil {
				fmt.Println("failed to convert pokemon to pgdata:", err)
			}
			_, err = queries.CreatePokemon(ctx, pkmn)
		}

		url = apiResponse.Next
		if url == "" {
			break
		}
	}

	pool.Close()
}

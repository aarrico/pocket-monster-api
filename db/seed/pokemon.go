package seed

import (
	"encoding/json"
	"fmt"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/aarrico/pocket-monster-api/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
	url2 "net/url"
	"strconv"
	"strings"
)

func populateTypes(types []Types, dbParams *db.CreatePokemonParams) {
	for _, typeData := range types {
		typeId, _ := queries.GetTypeByName(ctx, strings.ToLower(typeData.Type.Name))
		if typeData.Slot == 1 {
			dbParams.PrimaryType = typeId
		} else {
			dbParams.SecondaryType = pgtype.Int4{Int32: typeId, Valid: true}
		}
	}
}

func populateBaseStats(stats []BaseStats, dbParams *db.CreatePokemonParams) {
	for _, statData := range stats {
		statVal := int32(statData.Value)
		switch statData.Stat.Name {
		case "attack":
			dbParams.BaseAttack = statVal
		case "defense":
			dbParams.BaseDefense = statVal
		case "special-attack":
			dbParams.BaseSpecialAttack = statVal
		case "special-defense":
			dbParams.BaseSpecialDefense = statVal
		case "speed":
			dbParams.BaseSpeed = statVal
		case "hp":
			dbParams.BaseHp = statVal
		}
	}
}

func getNationalDexOrder(speciesUrl string) int32 {
	u, _ := url2.Parse(speciesUrl)
	parts := strings.Split(u.Path, "/")
	dexOrderStr := parts[len(parts)-2]
	dexOrder, _ := strconv.Atoi(dexOrderStr)

	return int32(dexOrder)
}

func getPokemon(url string) (db.CreatePokemonParams, error) {
	body := utils.GetBodyFromUrl(url, true)

	var pkmn Pokemon
	var dbParams db.CreatePokemonParams
	if err := json.Unmarshal(body, &pkmn); err != nil {
		fmt.Println("error unmarshalling pokemon data:", err)
		return dbParams, err
	}

	dbParams.Name = pkmn.Name
	dbParams.Height = pgtype.Int4{Int32: pkmn.Height, Valid: true}
	dbParams.Weight = pgtype.Int4{Int32: pkmn.Weight, Valid: true}
	dbParams.NationalDexOrder = getNationalDexOrder(pkmn.Species.Url)
	dbParams.BaseExperience = pgtype.Int4{Int32: pkmn.BaseExperience, Valid: true}
	dbParams.SortOrder = pkmn.SortOrder
	dbParams.IsDefault = pkmn.IsDefault

	populateTypes(pkmn.Types, &dbParams)
	populateBaseStats(pkmn.BaseStats, &dbParams)

	return dbParams, nil
}

func PopulatePokemonTable() {
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
			id, err := queries.CreatePokemon(ctx, pkmn)
			pkmnNameToId[pkmn.Name] = id
		}

		url = apiResponse.Next
		if url == "" {
			break
		}
	}
}

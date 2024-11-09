package seed

import (
	"encoding/json"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/aarrico/pocket-monster-api/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
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

func populatePokemon(url string) {
	body := utils.GetBodyFromUrl(url, true)

	var pkmn Pokemon
	if err := json.Unmarshal(body, &pkmn); err != nil {
		log.Println("error unmarshalling pokemon data:", err)
		return
	}

	dbParams := db.CreatePokemonParams{
		Name:             pkmn.Name,
		Height:           pgtype.Int4{Int32: pkmn.Height, Valid: true},
		Weight:           pgtype.Int4{Int32: pkmn.Weight, Valid: true},
		NationalDexOrder: getNationalDexOrder(pkmn.Species.Url),
		BaseExperience:   pgtype.Int4{Int32: pkmn.BaseExperience, Valid: true},
		SortOrder:        pkmn.SortOrder,
		IsDefault:        pkmn.IsDefault,
	}

	populateTypes(pkmn.Types, &dbParams)
	populateBaseStats(pkmn.BaseStats, &dbParams)

	id, err := queries.CreatePokemon(ctx, dbParams)
	if err != nil {
		log.Println("error inserting pokemon:", err)
		return
	}

	pkmnNameToId[pkmn.Name] = id
}

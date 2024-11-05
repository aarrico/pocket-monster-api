package seed

import (
	"encoding/json"
	"fmt"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/aarrico/pocket-monster-api/internal/utils"
)

func setAbilityData(url string) {
	body := utils.GetBodyFromUrl(url, true)

	var ability Ability
	var dbParams db.CreateAbilityParams
	if err := json.Unmarshal(body, &ability); err != nil {
		fmt.Println("error unmarshalling ability data:", err)
	}

	dbParams.Name = ability.Name

	for _, effect := range ability.Entries {
		if effect.Language.Name == "en" {
			dbParams.Effect = effect.Effect
		}
	}

	abilityId, err := queries.CreateAbility(ctx, dbParams)
	if err != nil {
		fmt.Printf("failed to create ability %s:\n%s", dbParams.Name, err)
	}

	for _, pkmn := range ability.Pokemon {
		pkmnId, _ := queries.GetPokemonIdByName(ctx, pkmn.Pokemon.Name)
		pkmnAbilityParams := db.SetPokemonAbilityRelationParams{
			IsHidden:  pkmn.IsHidden,
			Slot:      int32(pkmn.Slot),
			PokemonID: pkmnId,
			AbilityID: abilityId,
		}
		queries.SetPokemonAbilityRelation(ctx, pkmnAbilityParams)
	}
}

func PopulateAbilityTable() {
	url := "https://pokeapi.co/api/v2/ability"

	for {
		body := utils.GetBodyFromUrl(url, true)

		var apiResponse ApiResp
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			fmt.Println("error unmarshalling api response:", err)
			break
		}

		for _, rawData := range apiResponse.Results {
			setAbilityData(rawData.Url)
		}

		url = apiResponse.Next
		if url == "" {
			break
		}
	}
}

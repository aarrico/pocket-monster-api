package seed

import (
	"encoding/json"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/aarrico/pocket-monster-api/internal/utils"
	"log"
)

func populateAbility(url string) {
	body := utils.GetBodyFromUrl(url, true)

	var ability Ability
	if err := json.Unmarshal(body, &ability); err != nil {
		log.Println("error unmarshalling ability data:", err)
		return
	}

	dbParams := db.CreateAbilityParams{Name: ability.Name}
	for _, effect := range ability.Entries {
		if effect.Language.Name == "en" {
			dbParams.Effect = effect.Effect
		}
	}

	abilityId, err := queries.CreateAbility(ctx, dbParams)
	if err != nil {
		log.Printf("failed to create ability %s:\n%s\n", dbParams.Name, err)
		return
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

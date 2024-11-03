package seed

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aarrico/pocket-monster-api/internal"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/jackc/pgx/v5/pgtype"
	"io"
	"net/http"
)

func seedPokemon(ctx context.Context, queries db.Queries) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/")
	if err != nil {
		fmt.Println("error getting all pokemon:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading pokemon body:", err)
		return
	}

	var pkmnData map[string]interface{}
	if err := json.Unmarshal(body, &pkmnData); err != nil {
		fmt.Println("error parsing pokemon:", err)
		return
	}

	height, _ := internal.ConvertToInt4(pkmnData["height"])
	weight, _ := internal.ConvertToInt4(pkmnData["weight"])
	nationalDexOrder, _ := internal.ConvertToInt4(pkmnData["id"])
	baseExp, _ := internal.ConvertToInt4(pkmnData["base_experience"])
	sortOrder, _ := internal.ConvertToInt4(pkmnData["order"])
	isDefault, _ := internal.ConvertToBool(pkmnData["is_default"])

	queries.CreatePokemon(ctx, db.CreatePokemonParams{
		Name:             pgtype.Text{String: pkmnData["name"].(string)},
		Height:           height,
		Weight:           weight,
		NationalDexOrder: nationalDexOrder,
		BaseExperience:   baseExp,
		IsDefault:        isDefault,
		SortOrder:        sortOrder,
	})
}

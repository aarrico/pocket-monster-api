package seed

import (
	"context"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/aarrico/pocket-monster-api/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type populateTableFunc func(string)

var pkmnNameToId map[string]pgtype.UUID
var ctx context.Context
var queries *db.Queries

func PopulateDb() {
	ctx = context.Background()
	pool := utils.ConnectToDb(ctx)
	queries = db.New(pool)

	pkmnNameToId = make(map[string]pgtype.UUID)

	populateTableFromApi("https://pokeapi.co/api/v2/pokemon", populatePokemon)
	populateTableFromApi("https://pokeapi.co/api/v2/ability", populateAbilities)

	pool.Close()
}

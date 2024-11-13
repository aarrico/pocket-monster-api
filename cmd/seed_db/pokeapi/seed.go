package pokeapi

import (
	"context"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
)

type cache struct {
	typeNameToId         map[string]int32
	pkmnNameToId         map[string]pgtype.UUID
	moveTargetNameToId   map[string]pgtype.UUID
	moveCategoryNameToId map[string]pgtype.UUID
}

func buildCaches(ctx context.Context, queries *db.Queries, populateCachesFromDb bool) cache {
	return cache{
		typeNameToId:         buildTypeCache(ctx, queries),
		pkmnNameToId:         buildPkmnCache(ctx, queries, populateCachesFromDb),
		moveTargetNameToId:   buildMoveTargetCache(ctx, queries, populateCachesFromDb),
		moveCategoryNameToId: buildMoveCategoryCache(ctx, queries, populateCachesFromDb),
	}
}

func buildTypeCache(ctx context.Context, queries *db.Queries) map[string]int32 {
	typeNameToId := make(map[string]int32)
	if types, err := queries.ListTypes(ctx); err != nil {
		log.Fatalf("could not get types from the db, shutting down:\n%s", err)
	} else {
		for _, t := range types {
			typeNameToId[t.Name] = t.ID
		}
	}

	return typeNameToId
}

var caches cache

func Seed(ctx context.Context, queries *db.Queries, seedAllTables bool) {
	caches = buildCaches(ctx, queries, !seedAllTables)

	if seedAllTables {
		populateTableFromApi(ctx, queries, "https://pokeapi.co/api/v2/pokemon", populatePokemon)
		populateTableFromApi(ctx, queries, "https://pokeapi.co/api/v2/ability", populateAbility)
		populateTableFromApi(ctx, queries, "https://pokeapi.co/api/v2/move-target", populateMoveTarget)
		populateTableFromApi(ctx, queries, "https://pokeapi.co/api/v2/move-category", populateMoveCategory)
		populateTableFromApi(ctx, queries, "https://pokeapi.co/api/v2/move", populateMove)
	}

}

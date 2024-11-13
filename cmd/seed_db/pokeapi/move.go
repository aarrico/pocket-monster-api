package pokeapi

import (
	"context"
	"encoding/json"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/aarrico/pocket-monster-api/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
)

func populateMoveTarget(ctx context.Context, queries *db.Queries, url string) {
	body := utils.GetBodyFromUrl(url, true)

	var moveTarget BasicInfo
	if err := json.Unmarshal(body, &moveTarget); err != nil {
		log.Println("error unmarshalling ability data:", err)
		return
	}

	dbParams := db.CreateMoveTargetParams{
		Name: moveTarget.Name,
	}

	for _, desc := range moveTarget.Descriptions {
		if desc.Language.Name == "en" {
			dbParams.Description = desc.Description
		}
	}

	targetId, err := queries.CreateMoveTarget(ctx, dbParams)
	if err != nil {
		log.Printf("failed to create move target %s:\n%s", dbParams.Name, err)
		return
	}

	caches.moveTargetNameToId[dbParams.Name] = targetId
}

func populateMoveCategory(ctx context.Context, queries *db.Queries, url string) {
	body := utils.GetBodyFromUrl(url, true)

	var moveCategory BasicInfo
	if err := json.Unmarshal(body, &moveCategory); err != nil {
		log.Println("error unmarshalling ability data:", err)
		return
	}

	dbParams := db.CreateMoveCategoryParams{
		Name: moveCategory.Name,
	}

	for _, desc := range moveCategory.Descriptions {
		if desc.Language.Name == "en" {
			dbParams.Description = desc.Description
		}
	}

	categoryId, err := queries.CreateMoveCategory(ctx, dbParams)
	if err != nil {
		log.Printf("failed to create move category %s:\n%s", dbParams.Name, err)
		return
	}

	caches.moveCategoryNameToId[dbParams.Name] = categoryId
}

func populateMove(ctx context.Context, queries *db.Queries, url string) {
	body := utils.GetBodyFromUrl(url, true)

	var move Move
	if err := json.Unmarshal(body, &move); err != nil {
		log.Println("error unmarshalling move data:", err)
		return
	}

	moveTypeId, err := queries.GetTypeByName(ctx, move.Type.Name)
	if err != nil {
		log.Printf("failed to get type id for move %s\ntype name %s:\n%s", move.Name, move.Type.Name, err)
		return
	}

	dbParams := db.CreateMoveParams{
		Name:          move.Name,
		Accuracy:      move.Accuracy,
		PowerPoints:   move.PowerPoints,
		Priority:      move.Priority,
		Power:         move.Power,
		DamageClass:   db.DamageClass(move.DamageClass.Name),
		TargetID:      caches.moveTargetNameToId[move.Target.Name],
		TypeID:        moveTypeId,
		Ailment:       db.MoveAilment(move.Meta.Ailment.Name),
		AilmentChance: move.Meta.AilmentChance,
		CategoryID:    caches.moveCategoryNameToId[move.Meta.Category.Name],
		MinHits:       move.Meta.MinHits,
		MaxHits:       move.Meta.MaxHits,
		MinTurns:      move.Meta.MinTurns,
		MaxTurns:      move.Meta.MaxTurns,
		Drain:         move.Meta.Drain,
		Healing:       move.Meta.Healing,
		CritRate:      move.Meta.CritRate,
		FlinchChance:  move.Meta.FlinchChance,
		StatChance:    move.Meta.StatChance,
	}

	for _, effect := range move.Effect {
		if effect.Language.Name == "en" {
			dbParams.Effect = effect.Effect
		}
	}

	_, err = queries.CreateMove(ctx, dbParams)
	if err != nil {
		log.Printf("failed to add %s to db:\n%s", move.Name, err)
		return
	}
}

func buildMoveTargetCache(ctx context.Context, queries *db.Queries, populateCacheFromDb bool) map[string]pgtype.UUID {
	moveTargetNameToId := make(map[string]pgtype.UUID)

	if populateCacheFromDb {
		if moveTargets, err := queries.ListMoveTargets(ctx); err != nil {
			log.Fatalf("could not get move targets from the db, shutting down:\n%s", err)
		} else {
			for _, mt := range moveTargets {
				moveTargetNameToId[mt.Name] = mt.ID
			}
		}
	}

	return moveTargetNameToId
}

func buildMoveCategoryCache(ctx context.Context, queries *db.Queries, populateCacheFromDb bool) map[string]pgtype.UUID {
	moveCategoryNameToId := make(map[string]pgtype.UUID)
	if populateCacheFromDb {
		if moveCategories, err := queries.ListMoveCategories(ctx); err != nil {
			log.Fatalf("could not get move categories from the db, shutting down:\n%s", err)
		} else {
			for _, mc := range moveCategories {
				moveCategoryNameToId[mc.Name] = mc.ID
			}
		}
	}

	return moveCategoryNameToId
}

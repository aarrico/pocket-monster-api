package seed

import (
	"encoding/json"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/aarrico/pocket-monster-api/internal/utils"
	"log"
)

func populateMoveTarget(url string) {
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

	moveTargetNameToId[dbParams.Name] = targetId
}

func populateMoveCategory(url string) {
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

	moveCategoryNameToId[dbParams.Name] = categoryId
}

func populateMove(url string) {
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
		TargetID:      moveTargetNameToId[move.Target.Name],
		TypeID:        moveTypeId,
		Ailment:       db.MoveAilment(move.Meta.Ailment.Name),
		AilmentChance: move.Meta.AilmentChance,
		CategoryID:    moveCategoryNameToId[move.Meta.Category.Name],
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

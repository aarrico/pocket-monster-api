// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: move.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMove = `-- name: CreateMove :one
INSERT INTO move
(name,
 accuracy,
 power_points,
 priority,
 power,
 damage_class,
 effect,
 target_id,
 type_id,
 ailment,
 ailment_chance,
 category_id,
 min_hits,
 max_hits,
 min_turns,
 max_turns,
 drain,
 healing,
 crit_rate,
 flinch_chance,
 stat_chance)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
RETURNING id
`

type CreateMoveParams struct {
	Name          string
	Accuracy      int32
	PowerPoints   int32
	Priority      int32
	Power         int32
	DamageClass   DamageClass
	Effect        string
	TargetID      pgtype.UUID
	TypeID        int32
	Ailment       MoveAilment
	AilmentChance int32
	CategoryID    pgtype.UUID
	MinHits       int32
	MaxHits       int32
	MinTurns      int32
	MaxTurns      int32
	Drain         int32
	Healing       int32
	CritRate      int32
	FlinchChance  int32
	StatChance    int32
}

func (q *Queries) CreateMove(ctx context.Context, arg CreateMoveParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createMove,
		arg.Name,
		arg.Accuracy,
		arg.PowerPoints,
		arg.Priority,
		arg.Power,
		arg.DamageClass,
		arg.Effect,
		arg.TargetID,
		arg.TypeID,
		arg.Ailment,
		arg.AilmentChance,
		arg.CategoryID,
		arg.MinHits,
		arg.MaxHits,
		arg.MinTurns,
		arg.MaxTurns,
		arg.Drain,
		arg.Healing,
		arg.CritRate,
		arg.FlinchChance,
		arg.StatChance,
	)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const createMoveCategory = `-- name: CreateMoveCategory :one
INSERT INTO move_category
(name, description)
VALUES ($1, $2)
RETURNING id
`

type CreateMoveCategoryParams struct {
	Name        string
	Description string
}

func (q *Queries) CreateMoveCategory(ctx context.Context, arg CreateMoveCategoryParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createMoveCategory, arg.Name, arg.Description)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const createMoveTarget = `-- name: CreateMoveTarget :one
INSERT INTO move_target
(name, description)
VALUES ($1, $2)
RETURNING id
`

type CreateMoveTargetParams struct {
	Name        string
	Description string
}

func (q *Queries) CreateMoveTarget(ctx context.Context, arg CreateMoveTargetParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createMoveTarget, arg.Name, arg.Description)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const setPokemonMove = `-- name: SetPokemonMove :exec
INSERT INTO pokemon_move
(pokemon_id, move_id)
VALUES ($1, $2)
`

type SetPokemonMoveParams struct {
	PokemonID pgtype.UUID
	MoveID    pgtype.UUID
}

func (q *Queries) SetPokemonMove(ctx context.Context, arg SetPokemonMoveParams) error {
	_, err := q.db.Exec(ctx, setPokemonMove, arg.PokemonID, arg.MoveID)
	return err
}

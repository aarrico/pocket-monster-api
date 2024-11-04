// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Pokemon struct {
	ID               pgtype.UUID
	Name             pgtype.Text
	Height           pgtype.Int4
	Weight           pgtype.Int4
	NationalDexOrder pgtype.Int4
	BaseExperience   pgtype.Int4
	IsDefault        pgtype.Bool
	SortOrder        pgtype.Int4
}

type Type struct {
	ID   int32
	Name string
}

type TypeEffectiveness struct {
	AttackingTypeID int32
	DefendingTypeID int32
	Multiplier      pgtype.Numeric
}

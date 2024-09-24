package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Pokemon struct {
	bun.BaseModel `bun:"table:pokemon,alias:pkmn"`

	ID     uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Name   string    `bun:"name,notnull"`
	Height int       `bun:"height"`
	Weight int       `bun:"weight"`
	Order  int       `bun:"order"`
}

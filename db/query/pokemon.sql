-- name: GetPokemon :one
SELECT * FROM pokemon
WHERE id = $1;

-- name: ListPokemon :many
SELECT * from pokemon
ORDER BY sort_order;

-- name: CreatePokemon :one
INSERT INTO pokemon (
    name, height, weight, national_dex_order, base_experience, is_default, sort_order
) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;
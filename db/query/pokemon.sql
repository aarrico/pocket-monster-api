-- name: GetPokemon :one
SELECT *
FROM pokemon
WHERE id = $1;

-- name: ListPokemon :many
SELECT *
from pokemon
ORDER BY sort_order;

-- name: CreatePokemon :one
INSERT INTO pokemon (name,
                     height,
                     weight,
                     national_dex_order,
                     base_experience,
                     is_default,
                     sort_order,
                     primary_type,
                     secondary_type,
                     base_attack,
                     base_defense,
                     base_special_attack,
                     base_special_defense,
                     base_speed,
                     base_hp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
RETURNING id;
-- name: GetPokemon :one
SELECT * FROM pokemon
WHERE id = $1;

-- name: ListPokemon :many
SELECT * from pokemon
ORDER BY sort_order;

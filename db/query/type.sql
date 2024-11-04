-- name: GetTypeByName :one
SELECT id FROM type WHERE name = $1;

-- name: ListTypes :many
SELECT * FROM type;

-- name: GetMultiplierById :one
SELECT multiplier
FROM type_effectiveness
WHERE attacking_type_id = $1 AND defending_type_id = $2;

-- name: GetMultiplierByName :one
SELECT te.multiplier
FROM type_effectiveness as te
JOIN type at on te.attacking_type_id = at.id
JOIN type dt on te.defending_type_id = dt.id
WHERE at.name = $1 AND dt.name = $2;
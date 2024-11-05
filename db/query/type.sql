-- name: GetTypeByName :one
SELECT id
FROM type
WHERE name = $1;

-- name: ListTypes :many
SELECT *
FROM type;

-- name: GetMultiplierById :one
SELECT multiplier
FROM type_effectiveness
WHERE attacking_type_id = $1
  AND defending_type_id = $2;

-- name: GetMultiplierByName :one
SELECT te.multiplier
FROM type_effectiveness AS te
         INNER JOIN type at ON te.attacking_type_id = at.id
         INNER JOIN type dt ON te.defending_type_id = dt.id
WHERE at.name = $1
  AND dt.name = $2;
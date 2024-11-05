-- name: GetAbility :one
SELECT *
from ability
where id = $1;

-- name: GetAbilityByName :one
SELECT *
from ability
where name = $1;

-- name: GetAbilitiesForPokemon :many
SELECT a.id, a.name, a.effect, pa.slot, pa.is_hidden
FROM ability AS a
         INNER JOIN pokemon_ability pa ON a.id = pa.ability_id
WHERE pa.pokemon_id = $1
ORDER BY pa.slot;

-- name: GetAbilitiesForPokemonByDexOrder :many
SELECT a.id, a.name, a.effect, pa.slot, pa.is_hidden
FROM ability AS a
         INNER JOIN pokemon_ability pa ON a.id = pa.ability_id
         INNER JOIN pokemon p ON pa.pokemon_id = p.id
WHERE p.national_dex_order = $1
ORDER BY pa.slot;
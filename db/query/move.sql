-- name: CreateMove :one
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
RETURNING id;

-- name: CreateMoveTarget :one
INSERT INTO move_target
(name, description)
VALUES ($1, $2)
RETURNING id;

-- name: CreateMoveCategory :one
INSERT INTO move_category
(name, description)
VALUES ($1, $2)
RETURNING id;
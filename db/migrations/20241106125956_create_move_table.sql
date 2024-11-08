-- +goose Up
-- +goose StatementBegin
CREATE TYPE DAMAGE_CLASS AS ENUM ('physical', 'special', 'status');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TYPE MOVE_AILMENT AS ENUM
    (
        'none',
        'unknown',
        'paralysis',
        'sleep',
        'freeze',
        'burn',
        'poison',
        'confusion',
        'infatuation',
        'trap',
        'nightmare',
        'torment',
        'disable',
        'yawn',
        'heal-block',
        'no-type-immunity',
        'leech-seed',
        'embargo',
        'perish-song',
        'ingrain'
        );
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE move_target
(
    id          uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name        TEXT UNIQUE NOT NULL,
    description TEXT        NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE move_category
(
    id          uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name        TEXT UNIQUE NOT NULL,
    description TEXT        NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE move
(
    id             uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name           TEXT UNIQUE  NOT NULL,
    accuracy       INT          NOT NULL,
    power_points   INT          NOT NULL,
    priority       INT          NOT NULL,
    power          INT          NOT NULL,
    damage_class   DAMAGE_CLASS NOT NULL,
    effect         TEXT         NOT NULL,
    target_id      uuid         NOT NULL REFERENCES move_target (id),
    type_id        INT          NOT NULL REFERENCES type (id),
    ailment        MOVE_AILMENT NOT NULL,
    ailment_chance INT          NOT NULL,
    category_id    uuid         NOT NULL REFERENCES move_category (id),
    min_hits       INT          NOT NULL,
    max_hits       INT          NOT NULL,
    min_turns      INT          NOT NULL,
    max_turns      INT          NOT NULL,
    drain          INT          NOT NULL,
    healing        INT          NOT NULL,
    crit_rate      INT          NOT NULL,
    flinch_chance  INT          NOT NULL,
    stat_chance    INT          NOT NULL
);
CREATE INDEX idx_move_damage_class ON move (damage_class);
CREATE INDEX idx_move_ailment ON move (ailment);
CREATE INDEX idx_move_type ON move (type_id);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE pokemon_move
(
    pokemon_id uuid NOT NULL REFERENCES pokemon (id),
    move_id    uuid NOT NULL REFERENCES move (id),
    PRIMARY KEY (pokemon_id, move_id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE pokemon_move;
DROP TABLE move;
DROP TABLE move_target;
DROP TABLE move_category;
DROP TYPE DAMAGE_CLASS;
DROP TYPE MOVE_AILMENT;
-- +goose StatementEnd

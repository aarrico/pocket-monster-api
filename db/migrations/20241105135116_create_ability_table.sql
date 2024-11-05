-- +goose Up
-- +goose StatementBegin
CREATE TABLE ability
(
    id     uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name   TEXT UNIQUE NOT NULL,
    effect TEXT        NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE pokemon_ability
(
    pokemon_id uuid REFERENCES pokemon (id),
    ability_id uuid REFERENCES ability (id),
    slot       INT     NOT NULL,
    is_hidden  BOOLEAN NOT NULL,
    PRIMARY KEY (pokemon_id, ability_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pokemon_ability;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE ability;
-- +goose StatementEnd

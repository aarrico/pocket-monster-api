-- +goose Up
-- +goose StatementBegin
CREATE TYPE STATISTIC AS ENUM
    (
        'hp',
        'attack',
        'defense',
        'special-attack',
        'special-defense',
        'speed',
        'accuracy',
        'evasion'
        );
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE nature
(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(20) UNIQUE NOT NULL,
    stat_up   STATISTIC          NOT NULL,
    stat_down STATISTIC          NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE nature;
DROP TYPE STATISTIC;
-- +goose StatementEnd

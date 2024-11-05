-- +goose Up
-- +goose StatementBegin
CREATE TABLE type
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(20) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO type
    (name)
VALUES ('normal'),
       ('fire'),
       ('water'),
       ('grass'),
       ('electric'),
       ('ice'),
       ('fighting'),
       ('poison'),
       ('ground'),
       ('flying'),
       ('psychic'),
       ('bug'),
       ('rock'),
       ('ghost'),
       ('dragon'),
       ('dark'),
       ('steel'),
       ('fairy');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE type_effectiveness
(
    attacking_type_id INT REFERENCES type (id),
    defending_type_id INT REFERENCES type (id),
    multiplier        NUMERIC(3, 2) NOT NULL,
    PRIMARY KEY (attacking_type_id, defending_type_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO type_effectiveness
(attacking_type_id,
 defending_type_id,
 multiplier)
VALUES
    -- NORMAL
    (1, 1, 1.0),
    (1, 3, 1.0),
    (1, 2, 1.0),
    (1, 4, 1.0),
    (1, 11, 1.0),
    (1, 16, 1.0),
    (1, 14, 0.0),
    (1, 17, 0.5),
    (1, 18, 1.0),
    (1, 5, 1.0),
    (1, 8, 1.0),
    (1, 12, 1.0),
    (1, 7, 1.0),
    (1, 6, 1.0),
    (1, 9, 1.0),
    (1, 13, 0.5),
    (1, 10, 1.0),
    (1, 15, 1.0),

    -- WATER
    (3, 1, 1.0),
    (3, 3, 0.5),
    (3, 2, 2.0),
    (3, 4, 0.5),
    (3, 11, 1.0),
    (3, 16, 1.0),
    (3, 14, 1.0),
    (3, 17, 1.0),
    (3, 18, 1.0),
    (3, 5, 1.0),
    (3, 8, 1.0),
    (3, 12, 1.0),
    (3, 7, 1.0),
    (3, 6, 1.0),
    (3, 9, 2.0),
    (3, 13, 2.0),
    (3, 10, 1.0),
    (3, 15, 0.5),

    -- FIRE
    (2, 1, 1.0),
    (2, 3, 0.5),
    (2, 2, 0.5),
    (2, 4, 2.0),
    (2, 11, 1.0),
    (2, 16, 1.0),
    (2, 14, 1.0),
    (2, 17, 2.0),
    (2, 18, 1.0),
    (2, 5, 1.0),
    (2, 8, 1.0),
    (2, 12, 2.0),
    (2, 7, 1.0),
    (2, 6, 2.0),
    (2, 9, 1.0),
    (2, 13, 0.5),
    (2, 10, 1.0),
    (2, 15, 0.5),

    -- GRASS
    (4, 1, 1.0),
    (4, 3, 2.0),
    (4, 2, 0.5),
    (4, 4, 0.5),
    (4, 11, 1.0),
    (4, 16, 1.0),
    (4, 14, 1.0),
    (4, 17, 0.5),
    (4, 18, 1.0),
    (4, 5, 1.0),
    (4, 8, 0.5),
    (4, 12, 0.5),
    (4, 7, 1.0),
    (4, 6, 1.0),
    (4, 9, 2.0),
    (4, 13, 2.0),
    (4, 10, 0.5),
    (4, 15, 0.5),

    -- PSYCHIC
    (11, 1, 1.0),
    (11, 3, 1.0),
    (11, 2, 1.0),
    (11, 4, 1.0),
    (11, 11, 0.5),
    (11, 16, 0.0),
    (11, 14, 1.0),
    (11, 17, 0.5),
    (11, 18, 1.0),
    (11, 5, 1.0),
    (11, 8, 2.0),
    (11, 12, 1.0),
    (11, 7, 2.0),
    (11, 6, 1.0),
    (11, 9, 1.0),
    (11, 13, 1.0),
    (11, 10, 1.0),
    (11, 15, 1.0),

    -- --DARK
    (16, 1, 1.0),
    (16, 3, 1.0),
    (16, 2, 1.0),
    (16, 4, 1.0),
    (16, 11, 2.0),
    (16, 16, 0.5),
    (16, 14, 2.0),
    (16, 17, 1.0),
    (16, 18, 0.5),
    (16, 5, 1.0),
    (16, 8, 1.0),
    (16, 12, 1.0),
    (16, 7, 0.5),
    (16, 6, 1.0),
    (16, 9, 1.0),
    (16, 13, 1.0),
    (16, 10, 1.0),
    (16, 15, 1.0),

    -- GHOST
    (14, 1, 0.0),
    (14, 3, 1.0),
    (14, 2, 1.0),
    (14, 4, 1.0),
    (14, 11, 2.0),
    (14, 16, 0.5),
    (14, 14, 2.0),
    (14, 17, 1.0),
    (14, 18, 1.0),
    (14, 5, 1.0),
    (14, 8, 1.0),
    (14, 12, 1.0),
    (14, 7, 1.0),
    (14, 6, 1.0),
    (14, 9, 1.0),
    (14, 13, 1.0),
    (14, 10, 1.0),
    (14, 15, 1.0),

    -- STEEL
    (17, 1, 1.0),
    (17, 3, 0.5),
    (17, 2, 0.5),
    (17, 4, 1.0),
    (17, 11, 1.0),
    (17, 16, 1.0),
    (17, 14, 1.0),
    (17, 17, 0.5),
    (17, 18, 2.0),
    (17, 5, 0.5),
    (17, 8, 1.0),
    (17, 12, 1.0),
    (17, 7, 1.0),
    (17, 6, 2.0),
    (17, 9, 1.0),
    (17, 13, 2.0),
    (17, 10, 1.0),
    (17, 15, 1.0),

    -- FAIRY
    (18, 1, 1.0),
    (18, 3, 1.0),
    (18, 2, 0.5),
    (18, 4, 1.0),
    (18, 11, 1.0),
    (18, 16, 2.0),
    (18, 14, 1.0),
    (18, 17, 0.5),
    (18, 18, 1.0),
    (18, 5, 1.0),
    (18, 8, 0.5),
    (18, 12, 1.0),
    (18, 7, 2.0),
    (18, 6, 1.0),
    (18, 9, 1.0),
    (18, 13, 1.0),
    (18, 10, 1.0),
    (18, 15, 2.0),

    -- ELECTRIC
    (5, 1, 1.0),
    (5, 3, 2.0),
    (5, 2, 1.0),
    (5, 4, 0.5),
    (5, 11, 1.0),
    (5, 16, 1.0),
    (5, 14, 1.0),
    (5, 17, 1.0),
    (5, 18, 1.0),
    (5, 5, 0.5),
    (5, 8, 1.0),
    (5, 12, 1.0),
    (5, 7, 1.0),
    (5, 6, 1.0),
    (5, 9, 0.0),
    (5, 13, 1.0),
    (5, 10, 1.0),
    (5, 15, 0.5),

    -- POISON
    (8, 1, 1.0),
    (8, 3, 1.0),
    (8, 2, 1.0),
    (8, 4, 2.0),
    (8, 11, 1.0),
    (8, 16, 1.0),
    (8, 14, 0.5),
    (8, 17, 0.0),
    (8, 18, 2.0),
    (8, 5, 1.0),
    (8, 8, 0.5),
    (8, 12, 1.0),
    (8, 7, 1.0),
    (8, 6, 1.0),
    (8, 9, 0.5),
    (8, 13, 0.5),
    (8, 10, 1.0),
    (8, 15, 1.0),

    -- BUG
    (12, 1, 1.0),
    (12, 3, 1.0),
    (12, 2, 0.5),
    (12, 4, 2.0),
    (12, 11, 2.0),
    (12, 16, 2.0),
    (12, 14, 0.5),
    (12, 17, 0.5),
    (12, 18, 0.5),
    (12, 5, 1.0),
    (12, 8, 0.5),
    (12, 12, 1.0),
    (12, 7, 0.5),
    (12, 6, 1.0),
    (12, 9, 1.0),
    (12, 13, 1.0),
    (12, 10, 0.5),
    (12, 15, 1.0),

    -- FIGHTING
    (7, 1, 2.0),
    (7, 3, 1.0),
    (7, 2, 1.0),
    (7, 4, 1.0),
    (7, 11, 0.5),
    (7, 16, 2.0),
    (7, 14, 0.0),
    (7, 17, 2.0),
    (7, 18, 0.5),
    (7, 5, 1.0),
    (7, 8, 0.5),
    (7, 12, 0.5),
    (7, 7, 1.0),
    (7, 6, 2.0),
    (7, 9, 1.0),
    (7, 13, 2.0),
    (7, 10, 0.5),
    (7, 15, 1.0),

    -- ICE
    (6, 1, 1.0),
    (6, 3, 0.5),
    (6, 2, 0.5),
    (6, 4, 2.0),
    (6, 11, 1.0),
    (6, 16, 1.0),
    (6, 14, 1.0),
    (6, 17, 0.5),
    (6, 18, 1.0),
    (6, 5, 1.0),
    (6, 8, 1.0),
    (6, 12, 1.0),
    (6, 7, 1.0),
    (6, 6, 1.0),
    (6, 9, 2.0),
    (6, 13, 1.0),
    (6, 10, 2.0),
    (6, 15, 2.0),

    -- GROUND
    (9, 1, 1.0),
    (9, 3, 1.0),
    (9, 2, 2.0),
    (9, 4, 0.5),
    (9, 11, 1.0),
    (9, 16, 1.0),
    (9, 14, 1.0),
    (9, 17, 2.0),
    (9, 18, 1.0),
    (9, 5, 2.0),
    (9, 8, 2.0),
    (9, 12, 0.5),
    (9, 7, 1.0),
    (9, 6, 1.0),
    (9, 9, 1.0),
    (9, 13, 2.0),
    (9, 10, 0.0),
    (9, 15, 1.0),

    -- ROCK
    (13, 1, 1.0),
    (13, 3, 1.0),
    (13, 2, 2.0),
    (13, 4, 1.0),
    (13, 11, 1.0),
    (13, 16, 1.0),
    (13, 14, 1.0),
    (13, 17, 0.5),
    (13, 18, 1.0),
    (13, 5, 1.0),
    (13, 8, 1.0),
    (13, 12, 2.0),
    (13, 7, 0.5),
    (13, 6, 2.0),
    (13, 9, 0.5),
    (13, 13, 1.0),
    (13, 10, 2.0),
    (13, 15, 1.0),

    -- FLYING
    (10, 1, 1.0),
    (10, 3, 1.0),
    (10, 2, 1.0),
    (10, 4, 2.0),
    (10, 11, 1.0),
    (10, 16, 1.0),
    (10, 14, 1.0),
    (10, 17, 0.5),
    (10, 18, 1.0),
    (10, 5, 0.5),
    (10, 8, 1.0),
    (10, 12, 2.0),
    (10, 7, 2.0),
    (10, 6, 1.0),
    (10, 9, 1.0),
    (10, 13, 0.5),
    (10, 10, 1.0),
    (10, 15, 1.0),

    -- DRAGON
    (15, 1, 1.0),
    (15, 3, 1.0),
    (15, 2, 1.0),
    (15, 4, 1.0),
    (15, 11, 1.0),
    (15, 16, 1.0),
    (15, 14, 1.0),
    (15, 17, 0.5),
    (15, 18, 0.0),
    (15, 5, 1.0),
    (15, 8, 1.0),
    (15, 12, 1.0),
    (15, 7, 1.0),
    (15, 6, 1.0),
    (15, 9, 1.0),
    (15, 13, 1.0),
    (15, 10, 1.0),
    (15, 15, 2.0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE type
-- +goose StatementEnd

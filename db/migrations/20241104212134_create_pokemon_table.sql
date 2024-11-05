-- +goose Up
-- +goose StatementBegin
CREATE TABLE pokemon
(
    id                   uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name                 TEXT                     NOT NULL,
    height               INT,
    weight               INT,
    national_dex_order   INT                      NOT NULL,
    base_experience      INT,
    is_default           BOOLEAN                  NOT NULL,
    sort_order           INT                      NOT NULL,
    primary_type         INT REFERENCES type (id) NOT NULL,
    secondary_type       INT REFERENCES type (id),
    base_attack          INT                      NOT NULL,
    base_defense         INT                      NOT NULL,
    base_special_attack  INT                      NOT NULL,
    base_special_defense INT                      NOT NULL,
    base_speed           INT                      NOT NULL,
    base_hp              INT                      NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pokemon;
-- +goose StatementEnd

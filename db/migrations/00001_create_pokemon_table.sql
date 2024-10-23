-- +goose Up
-- +goose StatementBegin
CREATE TABLE pokemon (
                         id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
                         name text,
                         height int,
                         weight int,
                         national_dex_order int,
                         base_experience int,
                         is_default boolean,
                         sort_order int
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE pokemon;
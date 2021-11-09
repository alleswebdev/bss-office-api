-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists offices
(
    id          BIGSERIAL PRIMARY KEY,
    name        text not null,
    description text      default null,
    removed     BOOLEAN   default false,
    created     timestamp default NOW(),
    updated     timestamp default NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE offices;
-- +goose StatementEnd

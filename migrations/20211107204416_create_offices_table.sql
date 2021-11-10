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

CREATE INDEX offices_removed_idx ON offices(removed);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX offices_removed_idx;
DROP TABLE offices;
-- +goose StatementEnd

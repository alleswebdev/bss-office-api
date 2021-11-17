-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists offices
(
    id          BIGSERIAL,
    name        text not null,
    description text      default null,
    removed     BOOLEAN   default false,
    created_at     timestamp default NOW(),
    updated_at     timestamp default null
) PARTITION BY HASH (id);

CREATE INDEX offices_removed_idx ON offices(removed);

create table offices_0 partition of offices(primary key (id)) for values with (modulus 5, remainder 0);
create table offices_1 partition of offices(primary key (id)) for values with (modulus 5, remainder 1);
create table offices_2 partition of offices(primary key (id)) for values with (modulus 5, remainder 2);
create table offices_3 partition of offices(primary key (id)) for values with (modulus 5, remainder 3);
create table offices_4 partition of offices(primary key (id)) for values with (modulus 5, remainder 4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS offices_removed_idx;
DROP TABLE offices;
-- +goose StatementEnd

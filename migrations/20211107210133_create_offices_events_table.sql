-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists offices_events
(
    id        BIGSERIAL,
    office_id BIGSERIAL,
    type      smallint not null,
    status    smallint not null,
    payload   jsonb     default null,
    created_at   timestamp default NOW()
) PARTITION BY HASH (id);;

CREATE INDEX offices_events_service_id_idx ON offices_events(office_id);
CREATE INDEX offices_events_type_idx ON offices_events(type);
CREATE INDEX offices_events_status_idx ON offices_events(status);


create table offices_events_0 partition of offices_events(primary key (id)) for values with (modulus 5, remainder 0);
create table offices_events_1 partition of offices_events(primary key (id)) for values with (modulus 5, remainder 1);
create table offices_events_2 partition of offices_events(primary key (id)) for values with (modulus 5, remainder 2);
create table offices_events_3 partition of offices_events(primary key (id)) for values with (modulus 5, remainder 3);
create table offices_events_4 partition of offices_events(primary key (id)) for values with (modulus 5, remainder 4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX offices_events_service_id_idx;
DROP INDEX offices_events_type_idx;
DROP INDEX offices_events_status_idx;
DROP TABLE offices_events;
-- +goose StatementEnd

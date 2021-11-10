-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists offices_events
(
    id        BIGSERIAL PRIMARY KEY,
    office_id BIGSERIAL,
    type      smallint not null,
    status    smallint not null,
    payload   jsonb     default null,
    created   timestamp default NOW()
);

CREATE INDEX offices_events_service_id_idx ON offices_events(office_id);
CREATE INDEX offices_events_type_idx ON offices_events(type);
CREATE INDEX offices_events_status_idx ON offices_events(status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX offices_events_service_id_idx;
DROP INDEX offices_events_type_idx;
DROP INDEX offices_events_status_idx;
DROP TABLE offices_events;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists offices_events
(
    id        BIGSERIAL PRIMARY KEY,
    office_id BIGSERIAL,
    type      int8 not null,
    status    int8 not null,
    payload   jsonb     default null,
    created   timestamp default NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE offices_events;
-- +goose StatementEnd

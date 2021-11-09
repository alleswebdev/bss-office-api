-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists offices_events
(
    id        BIGSERIAL PRIMARY KEY,
    office_id BIGSERIAL,
    type      text not null,
    status    int8 not null,
    payload   jsonb     default null,
    created   timestamp default NOW(),
    CONSTRAINT fk_office_id
        FOREIGN KEY (office_id)
            REFERENCES offices (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE offices_events
    DROP CONSTRAINT "fk_office_id";
DROP TABLE offices_events;
-- +goose StatementEnd

-- +migrate Up
create table if not exists levels
(
    id         serial primary key,
    maps        jsonb,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
drop table levels;
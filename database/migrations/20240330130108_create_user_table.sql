-- +goose Up
create table if not exists users (
    id serial primary key,
    username varchar(255) not null default '',
    email varchar(255) not null default '',
    phone_number varchar(255) not null default '',
    salt varchar(255) not null default '',
    hashed_password varchar(255) not null default '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NULL
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

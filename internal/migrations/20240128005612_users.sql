-- +goose Up
CREATE TABLE users (
    id serial primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null
);

-- +goose Down
DROP TABLE users;

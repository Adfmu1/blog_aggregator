-- +goose Up
CREATE TABLE users (
	id int PRIMARY KEY,
	created_at timestamp not null,
	updated_at timestamp not null,
	name varchar not null
);

-- +goose Down
DROP TABLE users;

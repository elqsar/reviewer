-- +goose Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(128) UNIQUE NOT NULL,
  password VARCHAR(128) NOT NULL,
  email VARCHAR(128) NOT NULL
);

-- +goose Down
DROP TABLE users;
-- +goose Up
CREATE TABLE reviews (
  id SERIAL PRIMARY KEY,
  title VARCHAR(64) UNIQUE NOT NULL,
  body VARCHAR(128) NOT NULL,
  user_id INT NOT NULL REFERENCES users(id)
);


-- +goose Down
DROP TABLE reviews;

-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE expense(
  id SERIAL PRIMARY KEY,
  description VARCHAR,
  amount INT
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS expense;
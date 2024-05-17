-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  username VARCHAR(255) NOT NULL,
  password TEXT NOT NULL,
  name VARCHAR(255) NOT NULL,
  last_name VARCHAR(400),
  email VARCHAR(300)
)
-- +goose StatementBegin
-- +goose StatementEnd

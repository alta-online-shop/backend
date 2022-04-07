CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,

  name VARCHAR NOT NULL CHECK (name != ''),
  description TEXT,
  price BIGINT NOT NULL CHECK (price >= 0),

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
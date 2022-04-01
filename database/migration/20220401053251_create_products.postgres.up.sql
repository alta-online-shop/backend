CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,

  name VARCHAR,
  description TEXT,
  price BIGINT CHECK (price >= 0),

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
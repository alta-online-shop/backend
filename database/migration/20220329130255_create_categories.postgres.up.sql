CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,

  name VARCHAR NOT NULL CHECK (name != ''),
  description TEXT,

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,

  name VARCHAR,
  description TEXT,

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
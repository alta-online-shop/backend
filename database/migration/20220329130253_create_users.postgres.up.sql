CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,

  email VARCHAR NOT NULL CHECK (email != '') UNIQUE,
  password TEXT NOT NULL CHECK (password != ''),
  fullname VARCHAR NOT NULL CHECK (fullname != ''),

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
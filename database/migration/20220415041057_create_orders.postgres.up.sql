CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY,

  user_id INT NOT NULL,
  quantity INT NOT NULL,

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS ratings (
  id SERIAL PRIMARY KEY,

  user_id INT NOT NULL,
  product_id INT NOT NULL,
  count INT NOT NULL CHECK(count >= 0 AND count <= 5),

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
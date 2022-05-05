CREATE TABLE IF NOT EXISTS comments (
  id SERIAL PRIMARY KEY,

  user_id INT NOT NULL,
  product_id INT NOT NULL,
  comment_id INT,
  content TEXT NOT NULL CHECK(content != ''),

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
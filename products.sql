CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name VARCHAR,
  price DECIMAL,
  description VARCHAR,
  amount INTEGER
)
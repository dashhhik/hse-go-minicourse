CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  balance DECIMAL(10, 2) NOT NULL
);
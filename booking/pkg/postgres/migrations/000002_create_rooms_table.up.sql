CREATE TABLE rooms (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  level INTEGER NOT NULL,
  in_maintenance BOOLEAN NOT NULL DEFAULT FALSE
);
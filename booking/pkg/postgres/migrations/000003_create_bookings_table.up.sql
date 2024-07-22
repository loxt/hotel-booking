CREATE TABLE bookings (
  id SERIAL PRIMARY KEY,
  placed_at TIMESTAMP NOT NULL DEFAULT NOW(),
  start_at DATE NOT NULL,
  end_at DATE NOT NULL,
  room_id INTEGER NOT NULL REFERENCES rooms(id),
  guest_id INTEGER NOT NULL REFERENCES guests(id)
);
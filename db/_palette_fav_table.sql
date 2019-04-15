-- Drop the favorites table if it exists
DROP TABLE IF EXISTS paletteFavs;

-- Creates a table for favorite colors, UserID connects to ColorID to form a favorite
CREATE TABLE paletteFavs
(
  userId INT NOT NULL,
  paletteId INT NOT NULL,

  FOREIGN KEY
  (userId) REFERENCES users
  (id),

  FOREIGN KEY
  (paletteId) REFERENCES palettes
  (paletteID)
);
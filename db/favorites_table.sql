-- Drop the favorites table if it exists
DROP TABLE IF EXISTS favorites;

-- Creates a table for favorite colors, UserID connects to ColorID to form a favorite
CREATE TABLE favorites
(
  userId INT NOT NULL,
  userHash INT NOT NULL,
  colorHex CHAR(6) NOT NULL,

  FOREIGN KEY
  (userId) REFERENCES users
  (id),

  FOREIGN KEY
  (colorHex) REFERENCES colors
  (hex)
)
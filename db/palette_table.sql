-- Drop the pallette table if it exists
DROP TABLE IF EXISTS pallette;

-- Creates a table for pallets palletteName, CreatorID, and 3 Color Hashes
CREATE TABLE pallette
(
  creatorID INT NOT NULL,
  primaryColorHex CHAR(6) NOT NULL,
  secondaryColorHex CHAR(6) NOT NULL,
  tertiaryColorHex CHAR(6) NOT NULL,

  FOREIGN KEY
  (creatorID) REFERENCES users
  (id),

  FOREIGN KEY
  (primarycolorHex) REFERENCES colors
  (hex),

  FOREIGN KEY
  (secondaryColorHex) REFERENCES colors
  (hex),

  FOREIGN KEY
  (tertiaryColorHex) REFERENCES colors
  (hex)
);
-- Drop the pallette table if it exists
DROP TABLE IF EXISTS palettes;

-- Creates a table for pallets palletteName, CreatorID, and 3 Color Hashes
CREATE TABLE palettes
(
  creatorID INT NOT NULL,
  paletteID INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  paletteName VARCHAR
(255),
  primaryHex CHAR
(6) NOT NULL,
  secondaryHex CHAR
(6) NOT NULL,
  tertiaryHex CHAR
(6) NOT NULL,

  FOREIGN KEY
(creatorID) REFERENCES users
(id),

  FOREIGN KEY
(primaryHex) REFERENCES colors
(hex),

  FOREIGN KEY
(secondaryHex) REFERENCES colors
(hex),

  FOREIGN KEY
(tertiaryHex) REFERENCES colors
(hex)
);
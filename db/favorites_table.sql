-- Drop the favorites table if it exists
DROP TABLE IF EXISTS favorites;

-- Creates a table for favorite colors, UserID connects to ColorID to form a favorite
CREATE TABLE favorites
(
  userId INT NOT NULL,
  userHash BIGINT NOT NULL,
  colorHex CHAR(6) NOT NULL,

  FOREIGN KEY
  (userId) REFERENCES users
  (id),

  FOREIGN KEY
  (colorHex) REFERENCES colors
  (hex)
);

INSERT INTO favorites
  (userId, userHash, colorHex)
VALUES(1, 885536276, 'ff0000');

INSERT INTO favorites
  (userId, userHash, colorHex)
VALUES(2, 2949673445, 'ff0000')
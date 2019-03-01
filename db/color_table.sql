-- Drop colors table if it exists
DROP TABLE IF EXISTS colors;
-- Create the colors table
CREATE TABLE colors
(
  color VARCHAR(30),

  r TINYINT(1)
  UNSIGNED DEFAULT 0 ,

  g TINYINT
  (1) UNSIGNED  DEFAULT 0 ,

  b TINYINT
  (1)  UNSIGNED DEFAULT 0 ,

  a TINYINT
  (1)  UNSIGNED DEFAULT 100 ,

  hex VARCHAR
  (6) NOT NULL UNIQUE,

  creatorId INT NOT NULL,

  creatorHash INT NOT NULL,

  views INT DEFAULT 0,

  FOREIGN KEY
  (creatorId) REFERENCES users
  (id)

);

  -- Insert red, blue, and yellow into the db
  INSERT INTO colors
    (color,r, hex, creatorId, creatorHash)
  VALUES("True Red", 255, "ff0000", 1, 885536276);

  INSERT INTO colors
    (color,b, hex, creatorId, creatorHash)
  VALUES("True Blue", 255, "0000ff", 1, 885536276);

  INSERT INTO colors
    (color,g, hex, creatorId, creatorHash)
  VALUES("True Green", 255, "00ff00", 1, 885536276);

  INSERT INTO colors
    (color,r,g,b, hex, creatorId, creatorHash)
  VALUES("Semi Pink", 220, 80, 125, "dc507d", 1, 885536276);



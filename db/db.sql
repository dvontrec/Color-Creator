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
  (7) NOT NULL UNIQUE,
  views INT DEFAULT 0
);


  -- Insert red, blue, and yellow into the db
  INSERT INTO colors
    (color,r, hex)
  VALUES("True Red", 255, "#ff0000");
  INSERT INTO colors
    (color,b, hex)
  VALUES("True Blue", 255, "#0000ff");
  INSERT INTO colors
    (color,g, hex)
  VALUES("True Green", 255, "#00ff00");

  INSERT INTO colors
    (color,r,g,b, hex)
  VALUES("Semi Pink", 220, 80, 125, "#dc507d");
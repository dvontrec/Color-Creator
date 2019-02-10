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
  (1)  UNSIGNED DEFAULT 255 ,
  
  views INT DEFAULT 0
);


  -- Insert red, blue, and yellow into the db
  INSERT INTO colors
    (color,r)
  VALUES("red", 255);
  INSERT INTO colors
    (color,b)
  VALUES("blue", 255);
  INSERT INTO colors
    (color,g)
  VALUES("green", 255);
-- Drop colors table if it exists
DROP TABLE IF EXISTS colors;
-- Create the colors table
CREATE TABLE colors
(
  color VARCHAR(30),
  r VARCHAR(3) DEFAULT "0",
  g VARCHAR(3) DEFAULT "0",
  b VARCHAR(3) DEFAULT "0",
  a VARCHAR(3) DEFAULT "0",
  views INT DEFAULT 0
);

-- Insert red, blue, and yellow into the db
INSERT INTO colors
  (color)
VALUES("red");
INSERT INTO colors
  (color)
VALUES("blue");
INSERT INTO colors
  (color)
VALUES("yellow");
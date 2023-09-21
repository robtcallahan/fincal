use fincal;

ALTER TABLE categories DROP COLUMN column_index;
ALTER TABLE categories ADD COLUMN column_index DOUBLE UNIQUE AFTER id;

ALTER TABLE categories DROP COLUMN letter;
ALTER TABLE categories DROP COLUMN is_category;

UPDATE `categories` SET column_index = 1 WHERE id = 11;
UPDATE `categories` SET column_index = 2 WHERE id = 12;
UPDATE `categories` SET column_index = 3 WHERE id = 13;
UPDATE `categories` SET column_index = 4 WHERE id = 14;
UPDATE `categories` SET column_index = 5 WHERE id = 15;
UPDATE `categories` SET column_index = 6 WHERE id = 16;
UPDATE `categories` SET column_index = 7 WHERE id = 17;
UPDATE `categories` SET column_index = 8 WHERE id = 18;
UPDATE `categories` SET column_index = 9 WHERE id = 19;
UPDATE `categories` SET column_index = 10 WHERE id = 22;
UPDATE `categories` SET column_index = 11 WHERE id = 23;
UPDATE `categories` SET column_index = 12 WHERE id = 25;
UPDATE `categories` SET column_index = 13 WHERE id = 28;
UPDATE `categories` SET column_index = 14 WHERE id = 29;
UPDATE `categories` SET column_index = 15 where id = 30;
UPDATE `categories` SET column_index = 16 where id = 31;
UPDATE `categories` SET column_index = 17 WHERE id = 40;
UPDATE `categories` SET column_index = 18 WHERE id = 41;
UPDATE `categories` SET column_index = 19 where id = 42;
UPDATE `categories` SET column_index = 20 where id = 46;
UPDATE `categories` SET column_index = 21 where id = 47;
UPDATE `categories` SET column_index = 22 where id = 54;
UPDATE `categories` SET column_index = 23 where id = 58;
UPDATE `categories` SET column_index = 24 where id = 59;
UPDATE `categories` SET column_index = 25 WHERE id = 61;
UPDATE `categories` SET column_index = 26 WHERE id = 67;

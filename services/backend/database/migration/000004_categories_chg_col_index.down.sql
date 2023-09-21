use fincal;

ALTER TABLE categories DROP COLUMN column_index;
ALTER TABLE categories ADD COLUMN column_index INT UNIQUE AFTER budget_period;
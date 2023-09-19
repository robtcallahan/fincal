use fincal;

ALTER TABLE `categories` CHANGE COLUMN `budget_group` `budget_group` ENUM('Discretionary', 'Non-Discretionary', 'Charity', 'Loans', 'Savings') DEFAULT 'Discretionary';

ALTER TABLE categories DROP INDEX name_UNIQUE;
ALTER TABLE categories ADD UNIQUE INDEX name_UNIQUE ((CASE WHEN deleted_at IS NULL THEN name END));

ALTER TABLE categories DROP INDEX column_index;
ALTER TABLE categories ADD UNIQUE INDEX column_index ((CASE WHEN deleted_at IS NULL THEN column_index END));

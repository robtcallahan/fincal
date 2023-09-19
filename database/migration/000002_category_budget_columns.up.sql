use `fincal`;

ALTER TABLE `categories` ADD COLUMN `budget_group` ENUM('Discretionary', 'Non-Discretionary', 'Loans', 'Savings') DEFAULT 'Discretionary' AFTER `color`;
ALTER TABLE `categories` ADD COLUMN `budget_amount` DECIMAL(8,2) DEFAULT 0.00 AFTER `budget_group`;
ALTER TABLE `categories` ADD COLUMN `budget_period` ENUM('Weekly', 'Monthly', 'Every 2 Weeks', 'Twice Monthly', 'Yearly') DEFAULT 'Monthly' AFTER `budget_amount`;

ALTER TABLE `categories` ADD COLUMN `weekly_amount` decimal (8,2) DEFAULT NULL AFTER `budget_period`;
ALTER TABLE `categories` ADD COLUMN `monthly_amount` decimal (8,2) DEFAULT NULL AFTER `weekly_amount`;
ALTER TABLE `categories` ADD COLUMN `twice_monthly_amount` decimal (8,2) DEFAULT NULL AFTER `monthly_amount`;
ALTER TABLE `categories` ADD COLUMN `every_two_weeks_amount` decimal (8,2) DEFAULT NULL AFTER `twice_monthly_amount`;
ALTER TABLE `categories` ADD COLUMN `yearly_amount` decimal (8,2) DEFAULT NULL AFTER `every_two_weeks_amount`;

ALTER TABLE categories DROP COLUMN column_index;
ALTER TABLE categories ADD COLUMN column_index INT UNIQUE AFTER id;

UPDATE `categories` SET column_index = 1, categories.budget_amount = 14.00, categories.budget_period = 'Weekly' where id = 11;
UPDATE `categories` SET column_index = 2, categories.budget_amount = 10.00, categories.budget_period = 'Weekly' where id = 12;
UPDATE `categories` SET column_index = 3, categories.budget_amount = 25.00, categories.budget_period = 'Monthly' where id = 13;
UPDATE `categories` SET column_index = 4, categories.budget_amount = 150.00, categories.budget_period = 'Weekly' where id = 14;
UPDATE `categories` SET column_index = 5, categories.budget_amount = 10.00, categories.budget_period = 'Weekly' where id = 15;
UPDATE `categories` SET column_index = 6, categories.budget_amount = 90.00, categories.budget_period = 'Monthly' where id = 16;
UPDATE `categories` SET column_index = 7, categories.budget_amount = 109.77, categories.budget_period = 'Monthly' where id = 17;
UPDATE `categories` SET column_index = 8, categories.budget_amount = 171.28, categories.budget_period = 'Monthly' where id = 18;
UPDATE `categories` SET column_index = 9, categories.budget_amount = 78.03, categories.budget_period = 'Monthly' where id = 19;
UPDATE `categories` SET column_index = 10, categories.budget_amount = 105.00, categories.budget_period = 'Monthly' where id = 22;
UPDATE `categories` SET column_index = 11, categories.budget_amount = 10.00, categories.budget_period = 'Monthly' where id = 23;
UPDATE `categories` SET column_index = 12, categories.budget_amount = 90.00, categories.budget_period = 'Monthly' where id = 25;
UPDATE `categories` SET column_index = 13, categories.budget_amount = 1352.50, categories.budget_period = 'Monthly' where id = 28;
UPDATE `categories` SET column_index = 14, categories.budget_amount = 14.01, categories.budget_period = 'Monthly' where id = 29;
UPDATE `categories` SET column_index = 15 where id = 30;
UPDATE `categories` SET column_index = 16 where id = 31;
UPDATE `categories` SET column_index = 17, categories.budget_amount = 100.00, categories.budget_period = 'Monthly' where id = 40;
UPDATE `categories` SET column_index = 18, categories.budget_amount = 50.00, categories.budget_period = 'Monthly' where id = 41;
UPDATE `categories` SET column_index = 19 where id = 42;
UPDATE `categories` SET column_index = 20 where id = 46;
UPDATE `categories` SET column_index = 21 where id = 47;
UPDATE `categories` SET column_index = 22 where id = 54;
UPDATE `categories` SET column_index = 23 where id = 58;
UPDATE `categories` SET column_index = 24 where id = 59;
UPDATE `categories` SET column_index = 25, categories.budget_amount = 312.65, categories.budget_period = 'Monthly' where id = 61;
UPDATE `categories` SET column_index = 26, categories.budget_amount = 100.00, categories.budget_period = 'Monthly' where id = 67;

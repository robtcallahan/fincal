use `fincal`;

ALTER TABLE `categories` ADD COLUMN `budget_group` ENUM('Discretionary', 'Non-Discretionary', 'Loans', 'Savings') DEFAULT 'Discretionary' AFTER `color`;
ALTER TABLE `categories` ADD COLUMN `budget_amount` DECIMAL(8,2) DEFAULT NULL AFTER `budget_group`;
ALTER TABLE `categories` ADD COLUMN `budget_period` ENUM('Weekly', 'Monthly', 'Every 2 Weeks', 'Twice Monthly', 'Yearly') DEFAULT 'Monthly' AFTER `budget_amount`;

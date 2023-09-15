use fincal;

ALTER TABLE `categories` ADD COLUMN `weekly_amount` decimal (8,2) DEFAULT NULL AFTER `color`;
ALTER TABLE `categories` ADD COLUMN `monthly_amount` decimal (8,2) DEFAULT NULL AFTER `weekly_amount`;
ALTER TABLE `categories` ADD COLUMN `twice_monthly_amount` decimal (8,2) DEFAULT NULL AFTER `monthly_amount`;
ALTER TABLE `categories` ADD COLUMN `every_two_weeks_amount` decimal (8,2) DEFAULT NULL AFTER `twice_monthly_amount`;
ALTER TABLE `categories` ADD COLUMN `yearly_amount` decimal (8,2) DEFAULT NULL AFTER `every_two_weeks_amount`;


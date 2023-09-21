ALTER TABLE `categories` DROP COLUMN `weekly_amount`;
ALTER TABLE `categories` DROP COLUMN `monthly_amount`;
ALTER TABLE `categories` DROP COLUMN `twice_monthly_amount`;
ALTER TABLE `categories` DROP COLUMN `every_two_weeks_amount`;
ALTER TABLE `categories` DROP COLUMN `yearly_amount`;

ALTER TABLE `categories` ADD COLUMN `pay_period_contribution` DECIMAL(8,2) DEFAULT 0.00 AFTER `budget_period`;

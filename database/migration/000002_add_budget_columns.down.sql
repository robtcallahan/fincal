USE `fincal`;

ALTER TABLE `categories` DROP COLUMN `weekly_amount`;
ALTER TABLE `categories` DROP COLUMN `monthly_amount`;
ALTER TABLE `categories` DROP COLUMN `twice_monthly_amount`;
ALTER TABLE `categories` DROP COLUMN `every_two_weeks_amount`;
ALTER TABLE `categories` DROP COLUMN `yearly_amount`;

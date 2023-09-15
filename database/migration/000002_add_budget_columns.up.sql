use fincal;

DROP TABLE IF EXISTS `budget`;
CREATE TABLE `budget` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `categories_id` int NOT NULL,
                              `weekly_amount` decimal(8,2) DEFAULT NULL,
                              `monthly_amount` decimal(8,2) DEFAULT NULL,
                              `twice_monthly_amount` decimal(8,2) DEFAULT NULL,
                              `every_two_weeks_amount` decimal(8,2) DEFAULT NULL,
                              `yearly_amount` decimal(8,2) DEFAULT NULL,
                              `created_at` datetime(3) DEFAULT NULL,
                              `updated_at` datetime(3) DEFAULT NULL,
                              `deleted_at` datetime(3) DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `id_UNIQUE` (`id`),
                              CONSTRAINT `budget_categories_id_fk` FOREIGN KEY (`categories_id`) REFERENCES `categories` (`id`),
                              KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

USE `fincal`;

CREATE TABLE IF NOT EXISTS `users`
(
    `id`         int          NOT NULL AUTO_INCREMENT,
    `email`      varchar(120) NOT NULL,
    `password`   varchar(40)  NOT NULL,
    `first_name` varchar(50)  NOT NULL,
    `last_name`  varchar(50)  NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;

INSERT INTO `users` VALUES (1, 'robtcallahan@gmail.com', 'xxxxxxxxxx', 'Rob', 'Callahan', '2023-09-19 10:09:00', NULL, NULL);

CREATE TABLE IF NOT EXISTS `income`
(
    `id`         int NOT NULL AUTO_INCREMENT,
    `users_id`   int NOT NULL,
    `income`     DECIMAL(8, 2)                                                          DEFAULT 0.00,
    `pay_period` ENUM ('Weekly', 'Monthly', 'Every 2 Weeks', 'Twice Monthly', 'Yearly') DEFAULT 'Monthly',
    `created_at` datetime(3)                                                            DEFAULT NULL,
    `updated_at` datetime(3)                                                            DEFAULT NULL,
    `deleted_at` datetime(3)                                                            DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_income_deleted_at` (`deleted_at`),
    CONSTRAINT `income_users_id_fk` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;

INSERT INTO `income` VALUES(1, 1, 600, 'Weekly', '2023-09-19 10:09:00', NULL, NULL);
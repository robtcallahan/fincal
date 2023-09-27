USE `fincal`;

CREATE TABLE IF NOT EXISTS `users`
(
    `id`         INT          NOT NULL AUTO_INCREMENT,
    `username`   VARCHAR(16)  NOT NULL,
    `first_name` VARCHAR(50)  NOT NULL,
    `last_name`  VARCHAR(50)  NOT NULL,
    `email`      VARCHAR(120) NOT NULL,
    `created_at` DATETIME(3) DEFAULT NULL,
    `updated_at` DATETIME(3) DEFAULT NULL,
    `deleted_at` DATETIME(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;

INSERT INTO `users`
VALUES  (1, 'robtcallahan', 'Rob', 'Callahan', 'robtcallahan@gmail.com', '2023-09-26 01:21:11.0', '2023-09-26 01:21:11.0', NULL);

CREATE TABLE IF NOT EXISTS `secrets`
(
    `id`           INT         NOT NULL AUTO_INCREMENT,
    `users_id`     INT         NOT NULL,
    `password`     BINARY(32) NOT NULL,
    `token`        VARCHAR(40) DEFAULT NULL,
    `token_expiry` DATETIME(3) DEFAULT NULL,
    `created_at`   DATETIME(3) DEFAULT NULL,
    `updated_at`   DATETIME(3) DEFAULT NULL,
    `deleted_at`   DATETIME(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_secrets_deleted_at` (`deleted_at`),
    CONSTRAINT `secrets_users_id_fk` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `income`
(
    `id`         INT NOT NULL AUTO_INCREMENT,
    `users_id`   INT NOT NULL,
    `income`     DECIMAL(8, 2)                                                          DEFAULT 0.00,
    `pay_period` ENUM ('Weekly', 'Monthly', 'Every 2 Weeks', 'Twice Monthly', 'Yearly') DEFAULT 'Monthly',
    `created_at` DATETIME(3)                                                            DEFAULT NULL,
    `updated_at` DATETIME(3)                                                            DEFAULT NULL,
    `deleted_at` DATETIME(3)                                                            DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_income_deleted_at` (`deleted_at`),
    CONSTRAINT `income_users_id_fk` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;

# INSERT INTO `income`
# VALUES (1, 1, 600, 'Weekly', '2023-09-19 10:09:00', NULL, NULL);
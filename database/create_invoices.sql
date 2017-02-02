CREATE DATABASE IF NOT EXISTS invoices;

CREATE TABLE IF NOT EXISTS `invoices`.`invoices` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `document` VARCHAR(14) NOT NULL,
  `description` VARCHAR(256) NOT NULL,
  `amount` DECIMAL(16,2) NOT NULL,
  `reference_month` INT NOT NULL,
  `reference_year` INT NOT NULL,
  `created_at` DATETIME(0) NOT NULL,
  `is_active` TINYINT(1) NOT NULL,
  `desactive_at` DATETIME(0),
  PRIMARY KEY (`id`));

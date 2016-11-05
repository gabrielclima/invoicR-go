CREATE DATABASE IF NOT EXISTS database;

CREATE TABLE IF NOT EXISTS `database`.`invoices` (
  `document` VARCHAR(14) NOT NULL,
  `description` VARCHAR(256),
  `amount` DECIMAL(16,2) NOT NULL,
  `reference_month` INT NOT NULL,
  `reference_year` INT NOT NULL,
  `created_at` DATETIME(0) NOT NULL,
  `is_active` TINYINT(1) NOT NULL,
  `desactive_at` DATETIME(0),
  PRIMARY KEY (`document`));

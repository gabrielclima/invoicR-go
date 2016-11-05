CREATE DATABASE IF NOT EXISTS desafio_stone;

CREATE TABLE IF NOT EXISTS `desafio_stone`.`invoices` (
  `document` VARCHAR(14) NOT NULL,
  `description` VARCHAR(256),
  `amount` DECIMAL(16,2),
  `reference_month` INT,
  `reference_year` INT,
  `created_at` DATETIME(0) ,
  `is_active` TINYINT(1),
  `desactive_at` DATETIME(0),
  PRIMARY KEY (`document`));

CREATE TABLE population (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pref_code` VARCHAR(2) NOT NULL,
  `pref_name` VARCHAR(128) NOT NULL,
  `era_name` VARCHAR(128) NOT NULL,
  `era_year` INT NOT NULL,
  `year` INT NOT NULL,
  `population` INT NOT NULL,
  `male_population` INT NOT NULL,
  `female_population` INT NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

-- -----------------------------------------------------
-- Table `areas`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `areas` (
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `beans`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `beans` (
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `producing_area_id` CHAR(36) NOT NULL,
  `producer` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_beans_areas_idx` (`producing_area_id` ASC));


-- -----------------------------------------------------
-- Table `coffees`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `coffees` (
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `price` INT UNSIGNED NOT NULL,
  `cost` INT UNSIGNED NOT NULL,
  `style` TINYINT UNSIGNED NOT NULL,
  `bean_id` CHAR(36) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_coffees_beans1_idx` (`bean_id` ASC));

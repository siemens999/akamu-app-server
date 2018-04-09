-- MySQL Script generated by MySQL Workbench
-- Thu 08 Feb 2018 12:27:09 PM CET
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema akamu
-- -----------------------------------------------------
-- This is the akamu database schema.

-- ------------------------------------duel-----------------
-- Schema akamu
--
-- This is the akamu database schema.
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `akamu` DEFAULT CHARACTER SET utf8mb4 ;
USE `akamu` ;


-- -----------------------------------------------------
-- Table `akamu`.`university`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`university` (
 `iduniversity` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(127),
  `city` VARCHAR(45),
  `country` VARCHAR(45),
  PRIMARY KEY (`iduniversity`)
);

-- -----------------------------------------------------
-- Table `akamu`.`subject`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`subject` (
  `idsubject` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `shortform` VARCHAR(45) NOT NULL,
  `semester` INT NULL,
  `department` VARCHAR(45) NULL,
  `university` INT UNSIGNED NULL,
  `description` VARCHAR(511) NULL,
  `idmongo` VARCHAR(45) NULL,
  PRIMARY KEY (`idsubject`),
  UNIQUE INDEX `idsubject_UNIQUE` (`idsubject` ASC),
  INDEX `name_UNIQUE` (`name` ASC),
  INDEX `shortform_UNIQUE` (`shortform` ASC),
  UNIQUE INDEX `idmongo_UNIQUE` (`idmongo` ASC),
  CONSTRAINT `fk_subject_university`
  FOREIGN KEY (`university`)
  REFERENCES `akamu`.`university`(`iduniversity`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`maintainer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`maintainer` (
  `idmaintainer` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `login` VARCHAR(45) NOT NULL,
  `password` VARCHAR(100) NOT NULL,
  `level` INT NOT NULL,
  `name` VARCHAR(45) CHARACTER SET 'utf8mb4',
  `subject` INT UNSIGNED NULL,
  `university` INT UNSIGNED NOT NULL,
  `email` VARCHAR(45) NULL,
  `idmongo` VARCHAR(45) NULL,
  PRIMARY KEY (`idmaintainer`),
  UNIQUE INDEX `idmaintainer_UNIQUE` (`idmaintainer` ASC),
  UNIQUE INDEX `login_UNIQUE` (`login` ASC),
  INDEX `fk_maintainer_1_idx` (`subject` ASC),
  UNIQUE INDEX `idmongo_UNIQUE` (`idmongo` ASC),
  CONSTRAINT `fk_maintainer_subject`
  FOREIGN KEY (`subject`)
  REFERENCES `akamu`.`subject` (`idsubject`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_maintainer_university`
  FOREIGN KEY (`university`)
  REFERENCES `akamu`.`university`(`iduniversity`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`image`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`image` (
  `idimage` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `path` VARCHAR(128) NOT NULL,
  PRIMARY KEY (`idimage`),
  UNIQUE INDEX `idimage_UNIQUE` (`idimage` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`question`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`question` (
  `idquestion` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `author` INT UNSIGNED NULL,
  `text` TEXT NULL,
  `image` INT UNSIGNED NULL,
  `subject` INT UNSIGNED NULL,
  `idmongo` VARCHAR(45) NULL,
  `reviewed` TINYINT NOT NULL,
  `verified` TINYINT NOT NULL,
  `published` TINYINT NOT NULL,
  PRIMARY KEY (`idquestion`),
  UNIQUE INDEX `idquestion_UNIQUE` (`idquestion` ASC),
  INDEX `fk_question_author_idx` (`author` ASC),
  INDEX `fk_question_subject_idx` (`subject` ASC),
  CONSTRAINT `fk_question_author`
    FOREIGN KEY (`author`)
    REFERENCES `akamu`.`maintainer` (`idmaintainer`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_question_subject`
    FOREIGN KEY (`subject`)
    REFERENCES `akamu`.`subject` (`idsubject`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_question_image`
    FOREIGN KEY (`image`)
    REFERENCES `akamu`.`image` (`idimage`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`avatar`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`avatar` (
  `idavatar` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `image` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`idavatar`),
  UNIQUE INDEX `idavatar_UNIQUE` (`idavatar` ASC),
  INDEX `fk_avatar_image_idx` (`image` ASC),
  CONSTRAINT `fk_avatar_image`
    FOREIGN KEY (`image`)
    REFERENCES `akamu`.`image` (`idimage`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`title`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`title` (
  `idtitle` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `subject` INT UNSIGNED NULL,
  `unlock_score` INT NULL,
  `idmongo` VARCHAR(45) NULL,
  `unlock_win` INT NULL,
  PRIMARY KEY (`idtitle`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC),
  UNIQUE INDEX `idtitle_UNIQUE` (`idtitle` ASC),
  UNIQUE INDEX `idmongo_UNIQUE` (`idmongo` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`user` (
  `iduser` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `time_registered` DATETIME NOT NULL,
  `username` VARCHAR(45) NOT NULL,
  `password` VARCHAR(100) NOT NULL,
  `email` VARCHAR(45) NULL,
  `semester` INT UNSIGNED NULL,
  `experience` INT NOT NULL,
  `selected_avatar` INT UNSIGNED NULL,
  `selected_title` INT UNSIGNED NULL DEFAULT 0,
  `verified` TINYINT UNSIGNED NOT NULL DEFAULT 0,
  `university` VARCHAR(45) NULL,
  `idmongo` VARCHAR(45),
  PRIMARY KEY (`iduser`),
  UNIQUE INDEX `iduser_UNIQUE` (`iduser` ASC),
  UNIQUE INDEX `username_UNIQUE` (`username` ASC),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC),
  INDEX `fk_user_avatar_idx` (`selected_avatar` ASC),
  INDEX `fk_user_title_idx` (`selected_title` ASC),
  CONSTRAINT `fk_user_avatar`
    FOREIGN KEY (`selected_avatar`)
    REFERENCES `akamu`.`avatar` (`idavatar`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_user_title`
    FOREIGN KEY (`selected_title`)
    REFERENCES `akamu`.`title` (`idtitle`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`pool`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`pool` (
  `idpool` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `shortform` VARCHAR(45) NOT NULL,
  `description` VARCHAR(511) NULL,
  `image` INT UNSIGNED NOT NULL,
  `idmongo` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idpool`),
  UNIQUE INDEX `idpool_UNIQUE` (`idpool` ASC),
  UNIQUE INDEX `shortform_UNIQUE` (`shortform` ASC),
  CONSTRAINT `fk_pool_image`
  FOREIGN KEY (`image`)
  REFERENCES `akamu`.`image` (`idimage`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)

ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`pool_question`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`pool_question` (
  `pool` INT UNSIGNED NOT NULL,
  `question` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`pool`, `question`),
  INDEX `fk_pool_question_question_idx` (`question` ASC),
  CONSTRAINT `fk_pool_question_question`
    FOREIGN KEY (`question`)
    REFERENCES `akamu`.`question` (`idquestion`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_pool_question_pool`
    FOREIGN KEY (`pool`)
    REFERENCES `akamu`.`pool` (`idpool`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`classicduel`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`classicduel` (
  `idduel` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_challanger` INT UNSIGNED NOT NULL,
  `user_challanged` INT UNSIGNED NOT NULL,
  `question11` INT UNSIGNED NOT NULL,
  `question12` INT UNSIGNED NOT NULL,
  `question21` INT UNSIGNED NOT NULL,
  `question22` INT UNSIGNED NOT NULL,
  `pool1` INT UNSIGNED NOT NULL,
  `pool2` INT UNSIGNED NOT NULL,
  `status` INT(1) NOT NULL,
  `time_start` DATETIME NOT NULL,
  `time_registered`DATETIME NOT NULL,
  `time_changed` DATETIME NOT NULL,
  `time_end` DATETIME NULL,
  `score_challanger` INT NULL,
  `score_challanged` INT NULL,
  `winner` INT UNSIGNED NULL,
  `idmongo` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idduel`),
  UNIQUE INDEX `idduel_UNIQUE` (`idduel` ASC),
  INDEX `fk_duel_winner_idx` (`winner` ASC),
  INDEX `fk_duel_pool2_idx` (`pool2` ASC),
  INDEX `fk_duel_pool1_idx` (`pool1` ASC),
  INDEX `fk_duel_question22_idx` (`question22` ASC),
  INDEX `fk_duel_question21_idx` (`question21` ASC),
  INDEX `fk_duel_question12_idx` (`question12` ASC),
  INDEX `fk_duel_question11_idx` (`question11` ASC),
  INDEX `fk_duel_challanger_idx` (`user_challanger` ASC),
  INDEX `fk_duel_challanged_idx` (`user_challanged` ASC),
  UNIQUE INDEX `idmongo_UNIQUE` (`idmongo` ASC),
  CONSTRAINT `fk_duel_challanger`
    FOREIGN KEY (`user_challanger`)
    REFERENCES `akamu`.`user` (`iduser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_duel_challanged`
    FOREIGN KEY (`user_challanged`)
    REFERENCES `akamu`.`user` (`iduser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_duel_question11`
    FOREIGN KEY (`question11`)
    REFERENCES `akamu`.`question` (`idquestion`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_duel_question12`
    FOREIGN KEY (`question12`)
    REFERENCES `akamu`.`question` (`idquestion`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_duel_question21`
    FOREIGN KEY (`question21`)
    REFERENCES `akamu`.`question` (`idquestion`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_duel_question22`
    FOREIGN KEY (`question22`)
    REFERENCES `akamu`.`question` (`idquestion`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_duel_pool1`
    FOREIGN KEY (`pool1`)
    REFERENCES `akamu`.`pool` (`idpool`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_duel_pool2`
    FOREIGN KEY (`pool2`)
    REFERENCES `akamu`.`pool` (`idpool`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_duel_winner`
    FOREIGN KEY (`winner`)
    REFERENCES `akamu`.`user` (`iduser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`explanation`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`explanation` (
  `question` INT UNSIGNED NOT NULL,
  `text` TEXT NULL,
  `image` INT UNSIGNED NULL,
  PRIMARY KEY (`question`),
  UNIQUE INDEX `idexplanation_UNIQUE` (`question` ASC),
  CONSTRAINT `fk_explanation_question`
    FOREIGN KEY (`question`)
    REFERENCES `akamu`.`question` (`idquestion`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`answer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`answer` (
  `idanswer` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `question` INT UNSIGNED NOT NULL,
  `type` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`idanswer`),
  UNIQUE INDEX `isanswer_UNIQUE` (`question` ASC),
  CONSTRAINT `fk_answer_question`
    FOREIGN KEY (`question`)
    REFERENCES `akamu`.`question` (`idquestion`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_answer_answertype`
    FOREIGN KEY (`type`)
    REFERENCES `akamu`.`answertype` (`idanswertype`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`textinputanswer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`textinputanswer` (
  `idtextinputanswer` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `answer` INT UNSIGNED NOT NULL,
  `correct` TINYTEXT NOT NULL,
  PRIMARY KEY (`answer`),
  UNIQUE INDEX `idtextinputanswer_UNIQUE` (`idtextinputanswer` ASC),
  UNIQUE INDEX `answer` (`answer`),
  CONSTRAINT `fk_textinputanswer_answer`
    FOREIGN KEY (`answer`)
    REFERENCES `akamu`.`answer` (`idanswer`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`option`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`option` (
  `idoption` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `image` INT UNSIGNED NULL,
  `text` TEXT NULL,
  `correct` TINYTEXT NULL,
  `answer` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`idoption`),
  UNIQUE INDEX `idoption_UNIQUE` (`idoption` ASC),
  INDEX `fk_option_image_idx` (`image` ASC),
  CONSTRAINT `fk_option_image`
    FOREIGN KEY (`image`)
    REFERENCES `akamu`.`image` (`idimage`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_option_answer`
    FOREIGN KEY (`answer`)
    REFERENCES `akamu`.`answer` (`idanswer`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`report`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`report` (
  `idreport` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `question` INT UNSIGNED NOT NULL,
  `test` TINYTEXT NOT NULL,
  PRIMARY KEY (`idreport`, `question`),
  UNIQUE INDEX `idreport_UNIQUE` (`idreport` ASC),
  INDEX `fk_report_question_idx` (`question` ASC),
  CONSTRAINT `fk_report_question`
    FOREIGN KEY (`question`)
    REFERENCES `akamu`.`question` (`idquestion`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `akamu`.`answertype`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`answertype` (
  `idanswertype` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  PRIMARY KEY (`idanswertype`),
  UNIQUE INDEX `idanswertype_UNIQUE` (`idanswertype` ASC))
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `akamu`.`flashcard`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`flashcard` (
  `idflashcard` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `author` INT UNSIGNED NULL,
  `subject` INT UNSIGNED NULL,
  `creationdate` DATETIME NOT NULL,
  `lastmodified`DATETIME NOT NULL,
  `version` INT UNSIGNED NOT NULL,
  `fronttext` VARCHAR(1023) NOT NULL,
  `backtext` VARCHAR(1023) NOT NULL,
  `frontimage` INT UNSIGNED NULL,
  `backimage` INT UNSIGNED NULL,
  PRIMARY KEY (`idflashcard`),
  INDEX `author_UNIQUE` (`author` ASC),
  INDEX `subject_UNIQUE` (`subject` ASC),
  CONSTRAINT `fk_flashcard_subject`
  FOREIGN KEY (`subject`)
  REFERENCES `akamu`.`subject` (`idsubject`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_flashcard_author`
  FOREIGN KEY (`author`)
  REFERENCES `akamu`.`user` (`iduser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_flashcard_frontimage`
  FOREIGN KEY (`frontimage`)
  REFERENCES `akamu`.`image` (`idimage`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_flashcard_backimage`
  FOREIGN KEY (`backimage`)
  REFERENCES `akamu`.`image` (`idimage`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `akamu`.`traininglist`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `akamu`.`traininglist` (
  `author` INT UNSIGNED NOT NULL,
  `subject` INT UNSIGNED NOT NULL,
  `creationdate` DATETIME NOT NULL,
  `lastmodified`DATETIME NOT NULL,
  `version` INT UNSIGNED DEFAULT 0,
  `upvotes` INT UNSIGNED DEFAULT 0,
  `downvotes` INT UNSIGNED DEFAULT 0,
  PRIMARY KEY (`author`,`subject`),
  INDEX `author_UNIQUE` (`author` ASC, `subject` ASC),
  INDEX `subject_UNIQUE` (`subject` ASC, `author` ASC),
  CONSTRAINT `fk_traininglist_subject`
  FOREIGN KEY (`subject`)
  REFERENCES `akamu`.`subject` (`idsubject`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION,
  CONSTRAINT `fk_traininglist_author`
  FOREIGN KEY (`author`)
  REFERENCES `akamu`.`user` (`iduser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Insert test data
-- -----------------------------------------------------
INSERT INTO university (iduniversity, name, city, country) VALUES (1,"Test-University","Test-City","Test-Country");
INSERT INTO subject (idsubject, name, shortform, semester, department, university, description, idmongo) VALUES (1,"Test-Subject", "subject",1,"Test-Department",1,"This is a test Subject",1);
INSERT INTO image (idimage,path) VALUES (1,"path to the image");
INSERT INTO title (idtitle, name, subject, unlock_score, unlock_win) VALUES (1,"Test-Tilte",1,1,1);
INSERT INTO avatar (idavatar, image) VALUES (1,1);
INSERT INTO user (iduser, time_registered, username, password, email, semester, experience, selected_avatar, selected_title, verified, university) VALUES (1, "2015-09-18T00:00:00Z", "Test-User", "Test-Password", "test@email.com",1,1,1,1,1,TRUE);



INSERT INTO `answertype` (`name`) VALUES ('option');
INSERT INTO `answertype` (`name`) VALUES ('textinput');

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

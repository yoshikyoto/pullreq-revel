-- create database
CREATE SCHEMA IF NOT EXISTS `pullreq` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ;
USE `niconico-html5` ;

-- --------------------------------------------------------
-- create comments table
-- comments と複数形になっているのは、gormの仕様のため
-- --------------------------------------------------------
DROP TABLE IF EXISTS `pullreq`.`comments`;
CREATE TABLE `pullreq`.`comments` (
    id INT UNIQUE NOT NULL,
    owner VARCHAR(255),
    repo VARCHAR(255),
    pull_id VARCHAR(255),
    title VARCHAR(255),
    body VARCHAR(1024),
    user_name VARCHAR(255),
    file_path VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    status INT
) ENGINE = InnoDB;

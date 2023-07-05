CREATE DATABASE IF NOT EXISTS bookstore;

ALTER DATABASE bookstore
    DEFAULT CHARACTER SET utf8
    DEFAULT COLLATE utf8_general_ci;

GRANT ALL PRIVILEGES ON bookstore.* TO 'bookstore@%' IDENTIFIED BY 'bookstore';


DROP TABLE IF EXISTS account;

CREATE TABLE IF NOT EXISTS account
(
    id        INTEGER UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username  VARCHAR(50),
    password  VARCHAR(100),
    name      VARCHAR(50),
    avatar    VARCHAR(100),
    telephone VARCHAR(20),
    email     VARCHAR(100),
    location  VARCHAR(100),
    INDEX (username)
) engine = InnoDB;

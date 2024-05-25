-- init.sql

CREATE DATABASE IF NOT EXISTS Users;

USE Users;

CREATE TABLE IF NOT EXISTS user_info (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
);
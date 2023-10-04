CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

CREATE TABLE users(
    ID int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null,
    createdAt timestamp DEFAULT current_date()
)ENGINE=INNODB;
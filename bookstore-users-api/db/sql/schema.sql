CREATE DATABASE IF NOT EXISTS users_db;
USE users_db;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    first_name varchar(45),
    last_name varchar(45),
    email varchar(45) not null unique,
    password varchar(32) not null,
    status varchar(45) not null,
    date_created datetime not null
) ENGINE=INNODB;
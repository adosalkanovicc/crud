After cloning github repository, first thing you should do is creating a database called 'userdb'. After this, create a table 'users'.

CREATE DATABASE userdb;
USE userdb;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    age INT
);

After you have created databases and table, run the code by typing: go run main.go database.go


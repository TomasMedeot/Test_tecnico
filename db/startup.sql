DROP DATABASE IF EXISTS technical_test;
CREATE DATABASE technical_test;

USE technical_test;

CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) UNIQUE NOT NULL,
    password CHAR(15) NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE post (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    likes INT NOT NULL DEFAULT 0,
    content TEXT NOT NULL CHECK (LENGTH(content) >= 200),
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE repost (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (post_id) REFERENCES post(id)
);
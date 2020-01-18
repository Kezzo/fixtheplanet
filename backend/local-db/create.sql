DROP DATABASE IF EXISTS `fixtheplanet`;
CREATE DATABASE fixtheplanet;
USE fixtheplanet;
CREATE TABLE active_tables(
    table_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    table_name VARCHAR(255) NOT NULL,
    creation_time TIMESTAMP NOT NULL);
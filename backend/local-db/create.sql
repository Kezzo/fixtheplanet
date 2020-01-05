DROP DATABASE IF EXISTS `fixtheplanet`;
CREATE DATABASE fixtheplanet;
USE fixtheplanet;
CREATE TABLE issues(
    issue_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    repo VARCHAR(255) NOT NULL,
    issueNr INT NOT NULL,
    language VARCHAR(30) NULL);
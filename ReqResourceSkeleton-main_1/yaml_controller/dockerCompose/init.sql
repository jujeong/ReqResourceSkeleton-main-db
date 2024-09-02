CREATE DATABASE IF NOT EXISTS ketidatabase;

USE ketidatabase;

CREATE TABLE IF NOT EXISTS validity (
    file_name VARCHAR(255) PRIMARY KEY,
    request_sent BOOLEAN DEFAULT FALSE,
    request_sent_timestamp DATETIME,
    respond_received BOOLEAN DEFAULT FALSE,
    respond_received_timestamp DATETIME
);

CREATE TABLE IF NOT EXISTS yaml_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    filename VARCHAR(255) NOT NULL UNIQUE,
    encoded_content TEXT NOT NULL,
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
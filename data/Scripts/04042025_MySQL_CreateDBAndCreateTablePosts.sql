-- Create the database
CREATE DATABASE IF NOT EXISTS GoPostDB;

-- Use the database
USE GoPostDB;

-- Create the Posts table
CREATE TABLE IF NOT EXISTS Posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    userId INT NOT NULL,
    title VARCHAR(255),
    body TEXT
);

-- Insert sample posts
INSERT INTO Posts (userId, title, body) VALUES
(1, 'Hello World', 'This is the first post'),
(2, 'Go is Awesome', 'Go makes web APIs easy'),
(3, 'Learning SQL', 'Let’s query some data'),
(4, 'Untitled', 'This post has a body'),
(5, 'Just a Test', 'Another record for filtering');

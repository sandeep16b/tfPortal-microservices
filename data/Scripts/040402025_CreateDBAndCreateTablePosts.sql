-- Create database
CREATE DATABASE GoPostDB;
GO

-- Use the new database
USE GoPostDB;
GO

-- Create Posts table
CREATE TABLE Posts (
    id INT IDENTITY(1,1) PRIMARY KEY,
    userId INT NOT NULL,
    title NVARCHAR(100) NOT NULL,
    body NVARCHAR(MAX) NOT NULL
);
GO

-- Insert sample records
INSERT INTO Posts (userId, title, body) VALUES
(1, 'Hello World', 'This is the first post'),
(2, 'Go is Awesome', 'Go makes web APIs easy'),
(3, 'Learning SQL', 'Let’s query some data'),
(4, 'Untitled', 'This post has a body'),
(5, 'Just a Test', 'Another record for filtering');
GO

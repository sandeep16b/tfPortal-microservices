-- Use the database
USE GoPostDB;

-- Create the travel table
select * from testdb.postsCREATE TABLE `travel` (
  `name` varchar(255) DEFAULT NULL,
  `flightno` int NOT NULL,
  `departure` varchar(255) DEFAULT NULL,
  `arrival` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`flightno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
;


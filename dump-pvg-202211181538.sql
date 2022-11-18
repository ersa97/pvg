-- MySQL dump 10.13  Distrib 8.0.28, for Win64 (x86_64)
--
-- Host: localhost    Database: pvg
-- ------------------------------------------------------
-- Server version	8.0.28

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `firstname` varchar(100) NOT NULL,
  `lastname` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `phone` varchar(100) NOT NULL,
  `Email` varchar(100) NOT NULL,
  `Birthday` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (9,'ersa','ersa','arkhab','$2a$10$Z/OI7Haa1P/GgSynRKUvEe1RRhQGkRJIbkHbCZBWi6PjimReImQO2','081122334455','ersa123@gmail.com','1997-11-13 15:04:05'),(10,'ersa123','ersa','arkhab','$2a$10$tp7NrQd6WV0TELpbiQEjnedVnRsQN9rQWglHx8x349ET/1PSC8wLG','081122334455','ersa123@gmail.com','1997-11-13 15:04:05'),(11,'raven','ersa','raven','$2a$10$MHQApZi1JBQHD9G/HsrnweI3ZKSJ0vQ.1tWMRJzfSeCj3eqyC6JEm','081122334456','raven.ersa@gmail.com','1997-11-13 15:04:05'),(12,'raven1','ersa','raven','$2a$10$7xn8U0f6HXmInlk1PnOdjutN45fo1UfKtycAtuCrLc5HyUZCeV.pq','081122334457','ravenersa97@gmail.com','1997-11-13 15:04:05'),(13,'raven12','ersa','raven','$2a$10$w4nZ3xNvLVqCC3Yvs/0.z.K42KEEyq9P6Zv2heBTx3LE/7itRoeem','081122334454','raven.ersa97@gmail.com','1997-11-13 15:04:05'),(31,'safrizal99','safrizal','99','$2a$10$2bg6e/2vne.heHT7ajFy5OEHbvlggHw36Z1gMWHCD/VjL07glIVzK','081188112233','safrizal99@gmail.com','2003-06-06 15:04:05');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'pvg'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-11-18 15:38:14

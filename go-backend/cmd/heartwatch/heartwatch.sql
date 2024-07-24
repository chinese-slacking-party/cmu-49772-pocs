-- MariaDB dump 10.19  Distrib 10.5.23-MariaDB, for Linux (x86_64)
--
-- Host: cmu-49783-db.cleo2ooaa9zc.us-east-1.rds.amazonaws.com    Database: heart_monitoring
-- ------------------------------------------------------
-- Server version	8.0.35

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `daily_reports`
--

DROP TABLE IF EXISTS `daily_reports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `daily_reports` (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `min_heart_rate` int NOT NULL,
  `max_heart_rate` int NOT NULL,
  `health_prediction` text,
  `name` varchar(100) NOT NULL DEFAULT 'Anonymous',
  `total_dps` int NOT NULL DEFAULT '0',
  `hr_sum` double NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `daily_reports`
--

LOCK TABLES `daily_reports` WRITE;
/*!40000 ALTER TABLE `daily_reports` DISABLE KEYS */;
/*!40000 ALTER TABLE `daily_reports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `heart_rates`
--

DROP TABLE IF EXISTS `heart_rates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `heart_rates` (
  `id` int NOT NULL AUTO_INCREMENT,
  `timestamp` timestamp NOT NULL,
  `name` varchar(100) NOT NULL DEFAULT 'Anonymous',
  `age` int DEFAULT NULL,
  `gender` text,
  `heart_rate` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `heart_rates`
--

LOCK TABLES `heart_rates` WRITE;
/*!40000 ALTER TABLE `heart_rates` DISABLE KEYS */;
INSERT INTO `heart_rates` VALUES (1,'0000-00-00 00:00:00','Fred',25,'M',80),(2,'0000-00-00 00:00:00','Fred',25,'M',80),(3,'0000-00-00 00:00:00','Fred',25,'M',80),(4,'2024-07-24 01:31:06','Fred',25,'M',80),(5,'2024-07-24 01:31:11','Fred',25,'M',80),(6,'0000-00-00 00:00:00','test',0,'',12),(7,'0000-00-00 00:00:00','test',0,'',11.785),(8,'1970-01-01 00:16:40','test',0,'',11.785),(9,'1970-01-01 00:16:40','test',0,'',11.785),(10,'1970-01-01 00:33:20','test',0,'',49.783);
/*!40000 ALTER TABLE `heart_rates` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-07-24  2:14:20

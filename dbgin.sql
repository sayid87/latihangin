-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: May 16, 2024 at 05:05 PM
-- Server version: 8.0.30
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `dbgin`
--
CREATE DATABASE IF NOT EXISTS `dbgin` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
USE `dbgin`;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id_user` int NOT NULL,
  `nama` text NOT NULL,
  `email` varchar(255) NOT NULL,
  `telp` text NOT NULL,
  `jenis_kelamin` text NOT NULL,
  `hobi` text NOT NULL,
  `pekerjaan` text NOT NULL,
  `foto` text NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `warna_favorit` text NOT NULL,
  `alamat` text NOT NULL,
  `password` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id_user`, `nama`, `email`, `telp`, `jenis_kelamin`, `hobi`, `pekerjaan`, `foto`, `tanggal_lahir`, `warna_favorit`, `alamat`, `password`) VALUES
(2, 'Nano', 'anto@gmail.com', '088888888', 'Pria', 'Tidur', 'Pelayan', '6080697.png', '2008-11-12', '#00ffcc', 'Jakarta Timur', '$2a$10$z9Zn/RUUiyJfejtA/8ktguWiOPuXvnap4eaoo.tz7siYjvBmulf3G'),
(3, 'Nina', 'nina@gmail.com', '084554654', 'Wanita', 'Makan', 'Pramugari', 'WhatsApp Image 2022-08-15 at 9.04.08 AM.jpeg', '2008-11-01', '#00ff22', 'Jakarta Selatan', '$2a$10$fhegweMoDMKIdhg4NmecvOo9hi3RCoEhVA8CqgbWhnK8bgw.gn4iu');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id_user` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

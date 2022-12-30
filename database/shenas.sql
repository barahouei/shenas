-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Dec 30, 2022 at 06:18 PM
-- Server version: 10.9.4-MariaDB
-- PHP Version: 8.1.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `shenas`
--
CREATE DATABASE IF NOT EXISTS `shenas` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `shenas`;

-- --------------------------------------------------------

--
-- Table structure for table `answers`
--

CREATE TABLE `answers` (
  `aid` int(15) NOT NULL,
  `qid` int(15) NOT NULL,
  `answer` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `check_is_friend_answered`
--

CREATE TABLE `check_is_friend_answered` (
  `cid` int(11) NOT NULL,
  `user_telegram_id` bigint(20) NOT NULL,
  `friend_telegram_id` bigint(20) NOT NULL,
  `is_answered` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `check_is_user_answered`
--

CREATE TABLE `check_is_user_answered` (
  `cid` int(15) NOT NULL,
  `user_telegram_id` bigint(64) NOT NULL,
  `is_answered` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `friend_answers`
--

CREATE TABLE `friend_answers` (
  `faid` int(15) NOT NULL,
  `user_telegram_id` bigint(64) NOT NULL,
  `friend_telegram_id` bigint(64) NOT NULL,
  `qid` int(15) NOT NULL,
  `aid` int(15) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `questions`
--

CREATE TABLE `questions` (
  `qid` int(15) NOT NULL,
  `question` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `uid` int(15) NOT NULL,
  `user_telegram_id` bigint(64) NOT NULL,
  `username` varchar(64) DEFAULT NULL,
  `first_name` varchar(64) NOT NULL,
  `last_name` varchar(64) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `link` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `user_answers`
--

CREATE TABLE `user_answers` (
  `uaid` int(15) NOT NULL,
  `user_telegram_id` bigint(64) NOT NULL,
  `qid` int(15) NOT NULL,
  `aid` int(15) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `answers`
--
ALTER TABLE `answers`
  ADD PRIMARY KEY (`aid`);

--
-- Indexes for table `check_is_friend_answered`
--
ALTER TABLE `check_is_friend_answered`
  ADD PRIMARY KEY (`cid`,`user_telegram_id`);

--
-- Indexes for table `check_is_user_answered`
--
ALTER TABLE `check_is_user_answered`
  ADD PRIMARY KEY (`cid`,`user_telegram_id`);

--
-- Indexes for table `friend_answers`
--
ALTER TABLE `friend_answers`
  ADD PRIMARY KEY (`faid`,`user_telegram_id`);

--
-- Indexes for table `questions`
--
ALTER TABLE `questions`
  ADD PRIMARY KEY (`qid`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`uid`,`user_telegram_id`,`link`);

--
-- Indexes for table `user_answers`
--
ALTER TABLE `user_answers`
  ADD PRIMARY KEY (`uaid`,`user_telegram_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `answers`
--
ALTER TABLE `answers`
  MODIFY `aid` int(15) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `check_is_friend_answered`
--
ALTER TABLE `check_is_friend_answered`
  MODIFY `cid` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `check_is_user_answered`
--
ALTER TABLE `check_is_user_answered`
  MODIFY `cid` int(15) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `friend_answers`
--
ALTER TABLE `friend_answers`
  MODIFY `faid` int(15) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `questions`
--
ALTER TABLE `questions`
  MODIFY `qid` int(15) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `uid` int(15) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `user_answers`
--
ALTER TABLE `user_answers`
  MODIFY `uaid` int(15) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

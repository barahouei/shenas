-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jan 04, 2023 at 01:15 PM
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

--
-- Dumping data for table `answers`
--

INSERT INTO `answers` (`aid`, `qid`, `answer`) VALUES
(1, 1, 'سفید'),
(2, 1, 'سیاه'),
(3, 1, 'آبی'),
(4, 1, 'قرمز'),
(5, 1, 'سبز'),
(6, 1, 'بنفش'),
(7, 2, 'موفقیت'),
(8, 2, 'تحصیل'),
(9, 2, 'ثروت'),
(10, 2, 'آرامش'),
(11, 3, 'مرگ'),
(12, 3, 'تاریکی'),
(13, 3, 'عنکبوت'),
(14, 3, 'مکان‌های بسته'),
(15, 4, 'عدم موفقیت'),
(16, 4, 'بچه'),
(17, 4, 'ناگفته‌ها'),
(18, 4, 'انتخاب‌های اشتباه'),
(19, 5, 'موبایل'),
(20, 5, 'کتاب'),
(21, 5, 'خرید'),
(22, 5, 'کامپیوتر'),
(23, 5, 'سیگار'),
(24, 5, 'قلیان'),
(25, 6, 'ثروتمند'),
(26, 6, 'تحصیل‌کرده'),
(27, 6, 'ورزشکار'),
(28, 6, 'فهمیده و با درک'),
(29, 7, 'عصبانیت'),
(30, 7, 'بددهنی'),
(31, 7, 'بی‌معرفتی'),
(32, 7, 'بچه مسلکی'),
(37, 8, 'مهربانی'),
(38, 8, 'صداقت'),
(39, 8, 'با مرام بودن'),
(40, 8, 'دست و دلباز بودن'),
(41, 9, 'پیتزا'),
(42, 9, 'قرمه‌سبزی'),
(43, 9, 'سالاد'),
(44, 9, 'فرقی نداره'),
(45, 10, 'پولدار'),
(46, 10, 'معمولی'),
(47, 10, 'فرقی نداره'),
(51, 11, 'شجاعت'),
(52, 11, 'شکمو بودن'),
(53, 11, 'نظم'),
(54, 11, 'پشتکار'),
(55, 11, 'درک'),
(56, 12, 'چک و چونه'),
(57, 12, 'لمسی'),
(58, 12, 'خرید سریع'),
(59, 12, 'عدم علاقه به خرید'),
(60, 13, 'هلو'),
(61, 13, 'موز'),
(62, 13, 'پرتقال'),
(63, 13, 'آلبالو'),
(64, 13, 'سیب'),
(65, 13, 'فرقی نداره'),
(66, 14, 'رومنس'),
(67, 14, 'ماجراجویی'),
(68, 14, 'معمایی'),
(69, 14, 'اکشن'),
(70, 14, 'روانشناختی'),
(71, 15, 'طراحی'),
(72, 15, 'بازیگری'),
(73, 15, 'ورزشکار حرفه‌ای'),
(74, 15, 'نویسندگی'),
(75, 16, 'تماشای فیلم یا انیمیشن'),
(76, 16, 'ورزش'),
(77, 16, 'مطالعه'),
(78, 16, 'طبیعت‌گردی'),
(79, 17, 'ظاهر'),
(80, 17, 'ماشین'),
(81, 17, 'وسایل شخصی'),
(82, 1, 'هیچکدام'),
(83, 2, 'هیچکدام'),
(84, 3, 'هیچکدام'),
(85, 4, 'هیچکدام'),
(86, 5, 'هیچکدام'),
(87, 6, 'هیچکدام'),
(88, 7, 'هیچکدام'),
(89, 8, 'هیچکدام'),
(90, 9, 'هیچکدام'),
(91, 10, 'هیچکدام'),
(92, 11, 'هیچکدام'),
(93, 12, 'هیچکدام'),
(94, 13, 'هیچکدام'),
(95, 14, 'هیچکدام'),
(96, 15, 'هیچکدام'),
(97, 16, 'هیچکدام'),
(98, 17, 'هیچکدام'),
(99, 1, 'زرد'),
(100, 16, 'خواب'),
(101, 2, 'قدرت'),
(102, 2, 'ازدواج'),
(103, 3, 'سوسک'),
(104, 3, 'موش'),
(105, 3, 'ارتفاع'),
(106, 4, 'عزیز از دست رفته'),
(107, 4, 'تصمیمات گرفته نشده'),
(108, 4, 'عشق'),
(109, 4, 'گذشته'),
(110, 6, 'شاد و سرخوش'),
(111, 6, 'آرام و متین'),
(112, 6, 'مغرور'),
(113, 6, 'باوفا'),
(114, 7, 'پرحرفی'),
(115, 7, 'کم‌حرفی'),
(116, 7, 'بی‌احساسی'),
(117, 7, 'غرور'),
(118, 7, 'خنگ بودن'),
(119, 8, 'خوش‌رویی'),
(120, 8, 'خاکی بودن'),
(121, 8, 'خوش‌خنده'),
(122, 8, 'سنگ صبور'),
(123, 9, 'لازانیا'),
(124, 9, 'قیمه'),
(125, 9, 'کباب'),
(126, 9, 'ماهی'),
(127, 9, 'کله‌پاچه'),
(128, 17, 'آلودگی'),
(129, 17, 'تخت'),
(130, 17, 'غذای دهنی'),
(131, 12, 'گز کردن کل بازار'),
(132, 15, 'پزشکی'),
(133, 15, 'خلبانی'),
(134, 15, 'وکالت'),
(135, 15, 'نظامی'),
(136, 15, 'مهندسی'),
(137, 15, 'برنامه‌نویسی'),
(138, 16, 'بازی');

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

--
-- Dumping data for table `questions`
--

INSERT INTO `questions` (`qid`, `question`) VALUES
(1, 'رنگ مورد علاقه؟'),
(2, 'تو زندگی دنبال؟'),
(3, 'عمیق‌ترین ترس؟'),
(4, 'بزرگترین حسرت زندگی؟'),
(5, 'اعتیاد به؟'),
(6, 'مهم‌ترین ویژگی پارتنر ایده‌آل؟'),
(7, 'بدترین اخلاق؟'),
(8, 'بهترین اخلاق؟'),
(9, 'غذای مورد علاقه؟'),
(10, 'وضعیت مالی پارتنر ایده‌آل؟'),
(11, 'بارزترین ویژگی شخصی؟'),
(12, 'نوع خرید؟'),
(13, 'میوه مورد علاقه؟'),
(14, 'ژانر مورد علاقه؟'),
(15, 'شغل رویایی؟'),
(16, 'سرگرمی مورد علاقه؟'),
(17, 'حساس روی؟');

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
  MODIFY `aid` int(15) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=139;

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
  MODIFY `qid` int(15) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

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

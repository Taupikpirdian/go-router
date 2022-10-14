-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Oct 14, 2022 at 06:21 PM
-- Server version: 5.7.33
-- PHP Version: 7.4.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `salt_academy_exam_2`
--

-- --------------------------------------------------------

--
-- Table structure for table `article`
--

CREATE TABLE `article` (
  `id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  `date` timestamp NULL DEFAULT NULL,
  `banner` varchar(100) DEFAULT NULL,
  `author` varchar(100) DEFAULT NULL,
  `thumbs` varchar(100) DEFAULT NULL,
  `is_highlight` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `article`
--

INSERT INTO `article` (`id`, `category_id`, `date`, `banner`, `author`, `thumbs`, `is_highlight`) VALUES
(1, 1, '2022-10-14 16:22:09', 'banner.jpg', 'Taupik', 'banner.jpg', 1);

-- --------------------------------------------------------

--
-- Table structure for table `article_lang`
--

CREATE TABLE `article_lang` (
  `id` int(11) NOT NULL,
  `base_id` int(11) UNSIGNED NOT NULL,
  `slug` varchar(150) NOT NULL,
  `lang` varchar(5) NOT NULL,
  `title` varchar(150) NOT NULL,
  `text` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `article_lang`
--

INSERT INTO `article_lang` (`id`, `base_id`, `slug`, `lang`, `title`, `text`) VALUES
(1, 1, 'dampak-buruk-mie-instan', 'id', 'Dampak Buruk Mie Instan', 'Mie instan adalah makanan yang rasannya sangat enak. Tak heran bila makanan in banyak dicintai oleh masyarakat. Namun ternyata sering mengkonsumsi mie instan membawa dampak buruk bagi tubuh. Berdasarkan sejumlah hasil penelitian, terlalu banyak makan mie instan dapat meningkatkan risiko penyakit kanker, gangguan usus, ginjal hingga obesitas. Karena dampak buruk yang ditimbulkan, maka sebaiknya konsumsi mie instan dihindari. Sebisa mungkin ganti mie instan dengan makanan lain. Kalaupun ingin makan mie instan, sebaiknya beri tenggang waktu antara 2 – 3 hari. Jika sudah terbiasa dengan tenggang waktu tersebut, maka bisa diperpanjang menjadi 5 – 6 hari dan seterusnya. Hal ini dapat membantu siapapun mengurangi konsumsi mie instan dengan lebih baik.'),
(2, 1, 'the-bad-impact-of-instant-noodles', 'en', 'The Bad Impact of Instant Noodles', 'Instant noodles are foods that taste very good. No wonder this food is loved by many people. However, it turns out that often consuming instant noodles has a bad impact on the body. Based on a number of research results, eating too much instant noodles can increase the risk of cancer, intestinal disorders, kidneys to obesity. Because of the bad effects caused, it is better to avoid the consumption of instant noodles. As much as possible replace instant noodles with other foods. Even if you want to eat instant noodles, you should give a grace period of between 2-3 days. If you are used to the grace period, it can be extended to 5-6 days and so on. This can help anyone to reduce consumption of instant noodles better.');

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`) VALUES
(1);

-- --------------------------------------------------------

--
-- Table structure for table `categories_lang`
--

CREATE TABLE `categories_lang` (
  `id` int(11) NOT NULL,
  `base_id` int(10) UNSIGNED NOT NULL,
  `lang` varchar(5) NOT NULL,
  `slug` varchar(45) DEFAULT NULL,
  `name` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `categories_lang`
--

INSERT INTO `categories_lang` (`id`, `base_id`, `lang`, `slug`, `name`) VALUES
(1, 1, 'id', 'makanan', 'Makanan'),
(2, 1, 'en', 'food', 'Food');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `article`
--
ALTER TABLE `article`
  ADD PRIMARY KEY (`id`),
  ADD KEY `category_id` (`category_id`);

--
-- Indexes for table `article_lang`
--
ALTER TABLE `article_lang`
  ADD PRIMARY KEY (`id`),
  ADD KEY `base_id` (`base_id`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `categories_lang`
--
ALTER TABLE `categories_lang`
  ADD PRIMARY KEY (`id`),
  ADD KEY `base_id` (`base_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `article`
--
ALTER TABLE `article`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `article_lang`
--
ALTER TABLE `article_lang`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `categories_lang`
--
ALTER TABLE `categories_lang`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

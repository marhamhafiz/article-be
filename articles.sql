-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 11, 2022 at 06:44 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `articles`
--

-- --------------------------------------------------------

--
-- Table structure for table `posts`
--

CREATE TABLE `posts` (
  `id` int(11) NOT NULL,
  `title` varchar(200) NOT NULL,
  `content` text NOT NULL,
  `category` varchar(100) NOT NULL,
  `created_date` timestamp NULL DEFAULT NULL,
  `updated_date` timestamp NULL DEFAULT NULL,
  `status` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `posts`
--

INSERT INTO `posts` (`id`, `title`, `content`, `category`, `created_date`, `updated_date`, `status`) VALUES
(1, 'Minggu hujan terus udah musim hujan aja ya sekarang', 'mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini mengapa jadi begini ', 'Paper', '2022-09-09 16:11:08', '2022-09-11 16:13:41', 'draft'),
(2, 'Mencoba validasi form nih gimana yaaaaaaaaaa', 'validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni validasi form apa ni', 'Novel', '2022-09-09 16:12:46', '2022-09-11 13:16:33', 'draft'),
(3, 'Belajar Golang padahal malam minggu', 'Apa yang engkau cari anak muda Apa yang engkau cari anak muda Apa yang engkau cari anak muda Apa yang engkau cari anak muda Apa yang engkau cari anak muda Apa yang engkau cari anak muda Apa yang engkau cari anak muda Apa yang engkau cari anak muda', 'News', '2022-09-09 16:17:57', '2022-09-11 16:42:59', 'publish'),
(5, 'Udah minggu in cok pusing ga loooo wahahahhahahahah', 'Ngoding nih pusing aing Ngoding nih pusing aing  Ngoding nih pusing aing  Ngoding nih pusing aing  Ngoding nih pusing aing  Ngoding nih pusing aing Ngoding nih pusing aing Ngoding nih pusing aing Ngoding nih pusing aing Ngoding nih pusing aing Ngoding nih pusing aing Ngoding nih pusing aing Ngoding nih pusing aing Ngoding nih pusing aing Ngoding nih pusing aing  Ngoding nih pusing aing  Ngoding nih pusing aing Ngoding nih pusing aing ', 'Paper', '2022-09-10 10:56:16', '2022-09-11 12:37:15', 'draft'),
(143, 'asdasd nyoba ngubah yang ini dulu kali ya', 'mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy', 'News', '2022-09-11 11:53:15', '2022-09-11 16:12:27', 'trash'),
(146, 'Gw nyoba ngubah yang ini dulu kali ya', 'mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy mencoba terus sampai sukses cuy', 'News', '2022-09-11 11:59:22', '2022-09-11 16:09:47', 'publish');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=147;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

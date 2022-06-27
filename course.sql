-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 27 Jun 2022 pada 14.56
-- Versi server: 10.4.13-MariaDB
-- Versi PHP: 7.4.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `course`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `answers`
--

CREATE TABLE `answers` (
  `id` int(11) NOT NULL,
  `exercise_id` int(11) NOT NULL,
  `question_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `answer` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `answers`
--

INSERT INTO `answers` (`id`, `exercise_id`, `question_id`, `user_id`, `answer`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 5, 'b', '2022-06-17 12:55:44', '2022-06-17 12:55:44'),
(2, 1, 2, 5, 'c', '2022-06-17 12:55:44', '2022-06-17 12:55:44'),
(3, 1, 3, 5, 'a', '2022-06-17 12:55:44', '2022-06-17 12:55:44'),
(4, 1, 4, 5, 'c', '2022-06-17 12:55:44', '2022-06-17 12:55:44'),
(5, 1, 5, 5, 'd', '2022-06-17 12:55:44', '2022-06-17 12:55:44'),
(6, 1, 6, 5, 'b', '2022-06-17 12:55:44', '2022-06-17 12:55:44'),
(7, 1, 7, 5, 'd', '2022-06-17 13:01:35', '2022-06-17 13:01:35'),
(8, 1, 8, 5, 'c', '2022-06-17 13:01:35', '2022-06-17 13:01:35'),
(9, 1, 9, 5, 'b', '2022-06-17 13:01:35', '2022-06-17 13:01:35'),
(10, 1, 10, 5, 'b', '2022-06-17 13:01:35', '2022-06-17 13:01:35'),
(11, 2, 11, 5, 'd', '2022-06-26 21:16:34', '2022-06-26 21:16:34');

-- --------------------------------------------------------

--
-- Struktur dari tabel `exercises`
--

CREATE TABLE `exercises` (
  `id` int(11) NOT NULL,
  `title` text NOT NULL,
  `description` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `exercises`
--

INSERT INTO `exercises` (`id`, `title`, `description`) VALUES
(1, 'Olimpiade Matematika SMA', 'Olimpiade Matematika tingkat SMA Jawa Timur 2099'),
(2, 'Testing Ilmu Pengetahuan Alam', 'Test pengatahuan dasar ilmu pengetahuan alam');

-- --------------------------------------------------------

--
-- Struktur dari tabel `questions`
--

CREATE TABLE `questions` (
  `id` int(11) NOT NULL,
  `exercise_id` int(11) NOT NULL,
  `body` text NOT NULL,
  `option_a` text NOT NULL,
  `option_b` text NOT NULL,
  `option_c` text NOT NULL,
  `option_d` text NOT NULL,
  `correct_answer` text NOT NULL,
  `score` int(11) NOT NULL,
  `creator_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `questions`
--

INSERT INTO `questions` (`id`, `exercise_id`, `body`, `option_a`, `option_b`, `option_c`, `option_d`, `correct_answer`, `score`, `creator_id`, `created_at`, `updated_at`) VALUES
(1, 1, 'Berapa Jumlah hasil dari 1 + 1?', '0.2', '2', '2.2', '22', 'b', 10, 1, '2022-06-15 14:01:08', '2022-06-15 14:01:08'),
(2, 1, 'Berapa Jumlah hasil dari 2 + 2?', '4', '4.4', '44', '0.4', 'a', 10, 1, '2022-06-15 14:01:08', '2022-06-15 14:01:08'),
(3, 1, 'Berapa Jumlah hasil dari 1 x 1?', '0.1', '1', '1.1', '11', 'b', 10, 1, '2022-06-15 14:09:22', '2022-06-15 14:09:22'),
(4, 1, 'Berapa Jumlah hasil dari 3 x 3?', '999', '9', '9.9', '99', 'b', 10, 1, '2022-06-15 14:09:50', '2022-06-15 14:09:50'),
(5, 1, 'Berapa hasil dari 2 + 3?', '0.5', '0.55', '-5', '5', 'd', 10, 1, '2022-06-15 14:11:13', '2022-06-15 14:11:13'),
(6, 1, 'Berapa hasil dari 23 x 0.1?', '0.1', '0.23', '23', '2.3', 'd', 10, 1, '2022-06-15 14:12:07', '2022-06-15 14:12:07'),
(7, 1, 'Jika 3 - 2 = 1, berapakah hasil dari 3 + 1?', '4', '5', '6', '7', 'a', 10, 1, '2022-06-15 14:15:16', '2022-06-15 14:15:16'),
(8, 1, 'Jika 2 + 2 = 4, berapakah hasil dari 3 + 3?', '23', '33', '6', '5', 'c', 10, 1, '2022-06-15 14:15:43', '2022-06-15 14:15:43'),
(9, 1, 'Jika 10 + 1 = 11, berapakah hasil dari 30 x 1?', '31', '13', '4', '30', 'd', 10, 1, '2022-06-15 14:15:47', '2022-06-15 14:15:47'),
(10, 1, 'Berapa hasil dari 9 + 3?', '11', '12', '13', '14', 'b', 10, 1, '2022-06-15 14:15:50', '2022-06-15 14:15:50'),
(11, 2, 'ada berapa planet dalam galaxy bima', '5', '2', '7', '9', 'd', 10, 5, '0000-00-00 00:00:00', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `password` text NOT NULL,
  `no_hp` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `no_hp`, `created_at`, `updated_at`) VALUES
(1, 'super admin', 'admin@gmail.com', 'ini password', '08171122233333', '2022-06-15 13:02:09', '2022-06-15 13:02:09'),
(2, 'ahsan', 'ahsan@mail.com', '$2a$10$mM7/GAbcxBE1.Z2ALg83puE3Vcqn75UlAEexxa/xbIaQCmjb8PKoa', '', '2022-06-15 23:45:44', '2022-06-15 23:45:44'),
(3, 'ahsan', 'ahsan2@mail.com', '$2a$10$UOJAFVPbm4QtObtejd0fy.RduB5brNCqGx4Kv10XpAVxmYIhOT9Vi', '', '2022-06-16 19:21:49', '2022-06-16 19:21:49'),
(4, 'admin', 'admin@gmail.com', '$2a$10$DxK2iHPsGkb0FG74ipPesehErzNowe/V1Pi3jpQlOAclfjSYRdWhm', '123456', '2022-06-25 21:05:54', '2022-06-25 21:05:54'),
(5, 'gigi', 'gigi@gmail.com', '$2a$10$RtmcRfG79BuR90RjqsdMM.a79.tFkK/0HSSK0lXKcJSBi.ts9vgD6', '123456', '2022-06-25 21:35:13', '2022-06-25 21:35:13');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `answers`
--
ALTER TABLE `answers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `answers_user_id_IDX` (`user_id`,`question_id`) USING BTREE,
  ADD KEY `answers_FK` (`exercise_id`),
  ADD KEY `answers_FK_1` (`question_id`);

--
-- Indeks untuk tabel `exercises`
--
ALTER TABLE `exercises`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `questions`
--
ALTER TABLE `questions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `questions_FK` (`exercise_id`),
  ADD KEY `questions_FK_1` (`creator_id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `answers`
--
ALTER TABLE `answers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT untuk tabel `exercises`
--
ALTER TABLE `exercises`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `questions`
--
ALTER TABLE `questions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `answers`
--
ALTER TABLE `answers`
  ADD CONSTRAINT `answers_FK` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`),
  ADD CONSTRAINT `answers_FK_1` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`),
  ADD CONSTRAINT `answers_FK_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Ketidakleluasaan untuk tabel `questions`
--
ALTER TABLE `questions`
  ADD CONSTRAINT `questions_FK` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`),
  ADD CONSTRAINT `questions_FK_1` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

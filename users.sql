-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Hôte : localhost
-- Généré le : mar. 06 août 2024 à 14:05
-- Version du serveur : 8.0.35
-- Version de PHP : 8.2.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `usertwist`
--

-- --------------------------------------------------------

--
-- Structure de la table `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `username` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(80) COLLATE utf8mb4_general_ci NOT NULL,
  `token` varchar(10) COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `token`) VALUES
(1, 'user1', '$2a$12$A7T4cd3YA4PddrmtpKmep.RRcpmyZit3xb1UiNEmLA0.RRjvRyE06', '8Mb19OuQAJ'),
(2, 'user2', '$2a$12$0UTJeuzNije0N9TMWgh5re.CtE0FG.yy6.BgJi8YTfBTJeuzRGtDq', '1wVtw244Sg'),
(3, 'admin', '$2a$12$pOWQtOtdtERdjsJbreqcie4qhptcyiiWQab8OcQU/cnwYbbPs.sz.', 'a7P47jXfFM'),
(4, 'login', '$2a$12$.SOoCNvHsDJ9Mo6Da7dqReuC2XIy5goYvzCtwRYa4H1AnTCvn4B.m', 'KlMnOpQrSt'),
(5, 'Napoleon', '$2a$12$Vl9Rb3hRsIUVAS/2khes.eUCx9JOfm/l.wECmjd2V572j9vFM1zeq', '4dRf9uIOL3'),
(6, 'Cleopatra', '$2a$12$xizT1udcnH.l6M/Q4nSTc.6PgqoYT3ZjOKAoQ0xf4ixxxCR/KVWfS', 'hG7kLmN5Vx'),
(7, 'Einstein', '$2a$12$lh2CEDGRL5q/X0U3Cfyo.ugB3X1AggX8kPXsGgPWMdt7RyC/bIB0a', '2F6YzpH1rD'),
(8, 'Shakespeare', '$2a$12$yP1.v.Q4zRwqCcXIIQME/eHVTS74Ztp33d/Je5AJSv2FBklO3y5Uy', 'sQ8u7vPw3X'),
(9, 'Gandhi', '$2a$12$pot/ftJZTqLiS.2akwKsl.5dZxL5zNstW0x8wdxBlzjZgrENUBxoC', '9ZxNtJ6LkY'),
(10, 'Aristotle', '$2a$12$0cDBV0pUhTvPTsAZpbLx7e2tkYGfGr5ZH.XrI6NH312a1Rwo0TVEO', '0Bd3rFgH8q'),
(11, 'Caesar', '$2a$12$Xli0ul7xZfhUQQwoGf0GaeYlV3CMlM9d43uPokyUHyNbisEZrctZq', 'xU5mLvT2W1'),
(12, 'Newton', '$2a$12$uADgNu9WfBgj98Br5FpCiOAtzTOTFrE.VQB2KiHKmEs1YLPYcuDsO', 'q9Ef3KpJz4'),
(13, 'Curie', '$2a$12$malCiN6U8yl/RWMVHXWNGul7Oa2Z545o0rsPn1aB74/dkrR5hHOmu', '7YhF1nM5pX'),
(14, 'Tesla', '$2a$12$.hRhjCqGfCZgzBm7DywN7OGA.suw2hi/d6J5YjQpGDIhLkAPRNq3K', 'z6R4yLtQ3K'),
(15, 'Lincoln', '$2a$12$d4Jf/obS6ONy4FUOg77uJ.EAfbRx5EAiGkMHbNV6HwuRuIzIGLHA.', 'fW8k2PvLx9'),
(16, 'DaVinci', '$2a$12$ftiPAeTcenlzTq0D8ZI2EuMq27JdOu8BUxU4SXQi.VhvEG4iHs8iq', '5gQ0NsVy7L'),
(17, 'Columbus', '$2a$12$I9w.j/oHrduUJszYHOXiCO43rd43D3gO2J92uSr0U2CDfpNxB5HGS', '3tB1kGmH9X'),
(18, 'Plato', '$2a$12$RYVunt/MfhSj7ZdCNOlxLObFOfiRRtbaaXJBU3LD53vjJxhlOZfVu', 'w8V6pLdM2Y'),
(19, 'Beethoven', '$2a$12$ctVtEO/bWhFcYGHWVKgTGOTIneZ.A3H7jZRruB.fQy0GtH/JiJSmy', '0xP3yRtF9Q'),
(20, 'Mozart', '$2a$12$r/AZG0fj3jGPYTDYkn0jb.vP16a.gLc/D7o8TcWDzw8m2qOld1PvO', 'z4F6mNjH1W'),
(21, 'Galileo', '$2a$12$C9I3rIdayfutMYxtd0oVruWb/ssBZJ8rHJcQUU5rLMSI0iDe60qwG', '5K2tQpL7Vx'),
(22, 'Darwin', '$2a$12$KNeIJ9A8UU1biuKhiRop7uggToep0bIaDBex6tJLPUGFa12Ol7lAC', 'r8W1yMnF4X'),
(23, 'WrightBrothers', '$2a$12$dhGHAV/TTYcpPpdt.D34u.Ik.Rkn19Tee2xocvKxh8WOryMKlkWDS', '6zP5kJ3lTy'),
(24, 'Edison', '$2a$12$PxhL4rd8UwtO4WytFxbzw.uneolQG/kqR6u2zihGQ39V5Zqu.vjQG', 'f9B7vQ2pX4'),
(25, 'Armstrong', '$2a$12$wCeCHNb0kGF1/Kr2lqH0v.Da2805KbTEdINkz1l.RExdIlOoZJFb2', '0tL3mFy8Rk'),
(26, 'Washington', '$2a$12$57USJfq8d9GKREjkjOjd1ul1CiJRDvdAam5B/GBvpUBmgvq9YfZiS', '7P9r2WvTxY'),
(27, 'Roosevelt', '$2a$12$rDlflsGw3J7iz3Q5NRr6JuQ9Min1OOa/aC5QN669OUS.n3C.iMqhi', '1zF6yQ3mLp'),
(28, 'Freud', '$2a$12$dTojpaaPubrRTfRLSQo48OaymL58l1lgAC07XedARL04pIsWMEvjW', 'x4K8vNtJ9Y'),
(29, 'Marx', '$2a$12$UUnTktrd3I7lN/XNWRfYc.2mvGttbqnRn.uNlNdnrQXrMpD/Y83OC', '3W7kR5fMpQ'),
(30, 'Socrates', '$2a$12$gDj6lYBVE10eQVNM6KKH8uexRRQCpI3r42QaW3rsdr7pmJCWa5ke.', 'r6B9pFy2Lt'),
(31, 'Hemingway', '$2a$12$9FE3RQxNHBoVCSXnpGIfv.Ztga70qrklHZYVDn5xuUfeVzmj7JLcm', '0P3yW7nHxK'),
(32, 'Austen', '$2a$12$bupCj5.GMpTygoiSqOQFq.1.iplIfa1JOL4Z0WEH0G69tLkPaGY/m', 'f8L5kV9Q3x'),
(33, 'Kafka', '$2a$12$OZFwdXL/HC/bBWtyLZ/.F.mEdwi6Dz/nZI9jGrqyZa4DVDuspvZfy', '4tB2rJ6mPy'),
(34, 'Poe', '$2a$12$yfD.EuiJ5CdG9ounK67Ssufe7OIMtl21Phr5BKNFzPRBw3SsaCgbC', 'x1F9kP3vL6'),
(35, 'Hawking', '$2a$12$GhLVHres6H7VvsZpDljDFORX9roy.R8EFi5ikdBepbaZ86xeINC/2', '0W7yR4nF2p'),
(36, 'Copernicus', '$2a$12$pS918wmhhoELyKDMOabRgOmHyvG3YUNa9uaEVubOXX51GZU86btyy', 'k9L6pQ3fJx'),
(37, 'Lovelace', '$2a$12$xPpCAar92uJ6MXbSlxIw9OiL7DVVND7VamnzphmJbb4NgiDW0X5iW', '3tB8vM5rP0'),
(38, 'Franklin', '$2a$12$0nEVCAPL2YENaUNkMUTO3eLnimBFvx6D4yF7Snx2L.Bdtboj71kWq', 'f1W4pJ7kL2'),
(39, 'Tesla', '$2a$12$TmTtoakysAQ3NJx0ZHTKde.JVUzD6z4I1i7q3sEv2sTmFwzlzEzae', '6zP3yV9rF5'),
(40, 'Faraday', '$2a$12$iwg81SFEG6ezM6UTQ1mA8.e8zX6M3wjgstT1sVfS4.IQEX5nToIBi', '0tL1mN6yR8'),
(41, 'Pasteur', '$2a$12$OklAolflwSWHuom.djCuxOgt/dNpS5b6aSSfXg38Ysamv2ZPRZoI6', 'r4B9kP3vJ2'),
(42, 'Fleming', '$2a$12$rpYwiG6Vwsv9TW..idaRE.WKnurqNMG4tc6hS7q2TpphHZJGRhGIa', 'f8W5tL1pQ7'),
(43, 'Jobs', '$2a$12$XgD3G94HFMikJ48qoYU3OuMIe.t0CGyDNM0FDnkcfc.wTgmFDUpOW', '0zF2yN7mR6'),
(44, 'Gates', '$2a$12$9m0ww3jou2DYxxXahnKqtOaV7BYCf87B51ObJZuwQwtHsqYIklkSC', 'k3L9pJ5fV1'),
(45, 'Zuckerberg', '$2a$12$3PEwcLjopJ6mQrnP2B4VlufrMny2Ii1qkBHWOXpZi5RV9ulh.BJFa', 'x2B7vP6rF4'),
(46, 'Musk', '$2a$12$yK7CiR3hKnTeZ5.nSj8Yae1mMdsPczT3.9VDFFp4dzCh4qE4WCefq', '5tW3yL9kN0'),
(47, 'Bezos', '$2a$12$GEbgSgTLB7dKrH7pCVR2MexXII9zNPY2oaSoMrIbe0U8O9CgexUV.', 'r1F8pJ6vQ3'),
(48, 'Branson', '$2a$12$4Qh3zfM.7L/TZv.Iy3SaaOP4/C1z7.umiIndIvB2Yk6aKYUipDZku', '0tP7mL3yV5'),
(49, 'Buffett', '$2a$12$P.aQD.N6ovuXXZhu.hZC2.WaRoDtRwfuuey2W1tlF.dXKs3ZTnxjm', 'f9B2kN6rJ4'),
(50, 'Picasso', '$2a$12$LcG29tjBFVMPUrQmZtccceygimRT83Mx16KPFczpovcJpug4HR5Ne', 'x5W3pQ1tL8'),
(51, 'VanGogh', '$2a$12$JoS2uTGFDC7AL6kuFe1A5evLmSCqsEazxA5wBxB0zKcOWUFIqa7Ce', '4zP9yJ6vR2'),
(52, 'Dali', '$2a$12$2uKMXeZtkufShRU9qrSZ5ulmUZ9pQ8UFbZJUQs7W1.m/OoN4Vfw2i', 'k8L1mN7fV3'),
(53, 'Warhol', '$2a$12$mLUcPiLTxdGXImebch3YIuyulZ8u8UyMYxgecKljlFXfTzEdEeuMW', '3F9rJ4vP5Q'),
(54, 'Michelangelo', '$2a$12$1B578H.aAyOhQVhdHr0k/.Rw1X2Zlz/RmuNESD06r9Hj4uNKxmwwa', '0G7yP3nL8R'),
(55, 'leonardodaVinci', '$2a$12$TJoZbKTR1xQMwqeJk.wNF.SQt4QxuJqcv3ez.HHirGuwts8tJlvCm', '34f5g6h7j8'),
(56, 'queenElizabethI', '$2a$12$jdd5QguamTAt3WU8LhHm5egF1ZWSzDGNozqNTGSG60DWmj3sECHMa', 'k9l8m7n6o5'),
(57, 'sherlockHolmes', '$2a$12$eskZJaDDwS2WxvSRl7w4MuIMa/KIcmdEGEkHiSqLBfPLQtwgyuOb.', 'p1q2r3s4t5'),
(58, 'marieCurie', '$2a$12$YICS1zq5G/3iV7rtpZepcOEpYC09uuS/EC2W52OzmXWh78M7h8FHO', 'v6w7x8y9z0'),
(59, 'napoleonBonaparte', '$2a$12$cEcdRYyrHOaXpKYF4gTG1uGTlzbbmhvjg3G8PyTUac2.lzdnMrFMe', 'b1c2d3e4f5'),
(60, 'janeAusten', '$2a$12$HHCiIrq8Lu3YGOC1ikft0uZ.ccDB.qvOLwC698fgTsxC4tPSaCFcS', 'g6h7i8j9k0'),
(61, 'albertEinstein', '$2a$12$RcXYkOmE28/FffKbNRthy.HRkLSHWpILckhuwvYPvHt0/zxCUS6ZS', 'm1n2o3p4q5'),
(62, 'cleopatra', '$2a$12$JgW31R13Byus2kqhbjBngOs9WXQaNBpTcGyAWnNu.jfWmD3gouIiO', 's6t7u8v9w0'),
(63, 'michaelangelo', '$2a$12$1eG50rD5V7j.CVHhenEa5.La2UbuJpsPG4gVtEpd2xAG/9JG2wpMK', 'z1a2b3c4d5'),
(64, 'plato', '$2a$12$ta3SAo3CZRWgcqxTOWLrg.nuQCr2k3DyXtqSexPHizGFGm1VStfmu', 'f6g7h8i9j0'),
(65, 'juliusCaesar', '$2a$12$WZYiMA.TLXH2yIcHQBQNzupNS4v2U2RAdYMRcLyRY1VEKjFzQXwRy', 'k1l2m3n4o5'),
(66, 'alexanderTheGreat', '$2a$12$RMJI/GWKMZxhV8KB/DOHiOnqR6utoKmP56ZNxnkL21lrlf3zoRpFu', 'p6q7r8s9t0'),
(67, 'confucius', '$2a$12$Coi.jehJm5OdDgWcaHZ0rOhgQus6B/586SFhDaC7Uc7RXag/rKviC', 'v1w2x3y4z5'),
(68, 'williamShakespeare', '$2a$12$AQnmOYRLyuEqhHHOMCxs6uD0koKYHA58.bOizRBieXXYHv57Te2BC', 'b6c7d8e9f0'),
(69, 'galileoGalilei', '$2a$12$aHKM4xnkzpECoWPBBZLr3ek1Km90VOhdRjD7k9JieZd4z1pKuOuHO', 'g1h2i3j4k5');

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=70;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2020-11-20 15:35:27
-- 服务器版本： 8.0.22-0ubuntu0.20.10.2
-- PHP 版本： 7.4.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `dig`
--

-- --------------------------------------------------------

--
-- 表的结构 `article`
--

CREATE TABLE `article` (
  `articleId` bigint UNSIGNED NOT NULL COMMENT 'id',
  `type` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '类型',
  `subject` varchar(500) NOT NULL DEFAULT '' COMMENT '标题',
  `addTime` datetime NOT NULL DEFAULT '2020-11-19 00:00:00' COMMENT '添加时间',
  `isHot` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否热榜,0:非热，1，热',
  `url` varchar(500) NOT NULL DEFAULT '' COMMENT '链接地址',
  `domain` varchar(100) NOT NULL DEFAULT '' COMMENT '域',
  `userId` bigint UNSIGNED NOT NULL DEFAULT '0' COMMENT '推荐用户',
  `approvalStaffId` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '批准员工',
  `digSum` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '被顶的次数',
  `commentSum` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '被评论的次数',
  `isPublish` tinyint NOT NULL DEFAULT '0' COMMENT '0,未发布，1，已发布'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

--
-- 转存表中的数据 `article`
--
-- 转储表的索引
--

--
-- 表的索引 `article`
--
ALTER TABLE `article`
  ADD PRIMARY KEY (`articleId`),
  ADD KEY `isPublish` (`isPublish`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `article`
--
ALTER TABLE `article`
  MODIFY `articleId` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id', AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

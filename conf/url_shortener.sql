/*
Navicat MySQL Data Transfer

Source Server         : pwcong
Source Server Version : 100121
Source Host           : localhost:3306
Source Database       : url_shortener

Target Server Type    : MYSQL
Target Server Version : 100121
File Encoding         : 65001

Date: 2017-04-16 14:54:57
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for urls_long_short
-- ----------------------------
DROP TABLE IF EXISTS `urls_long_short`;
CREATE TABLE `urls_long_short` (
  `id` bigint(20) NOT NULL,
  `url_long` varchar(255) NOT NULL,
  `url_short` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_url_long` (`url_long`),
  UNIQUE KEY `unique_url_short` (`url_short`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
SET FOREIGN_KEY_CHECKS=1;

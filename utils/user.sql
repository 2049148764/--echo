/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : user

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2020-11-05 18:12:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'huliang', '123456', '17542162634');
INSERT INTO `user` VALUES ('2', '1', '1', '1');
INSERT INTO `user` VALUES ('3', '1', '1', '1');
INSERT INTO `user` VALUES ('4', '1', '1', '1');
INSERT INTO `user` VALUES ('5', '1', '1', '1');
INSERT INTO `user` VALUES ('6', '1', '1', '1');
INSERT INTO `user` VALUES ('7', '1', '1', '1');
INSERT INTO `user` VALUES ('8', 'huliang', '123456', '17542162634');
INSERT INTO `user` VALUES ('9', 'huliang', '123456', '17542162634');
INSERT INTO `user` VALUES ('14', 'huliang', '123456', '17542162634');
INSERT INTO `user` VALUES ('11', 'huliang', '123456', '17542162634');
INSERT INTO `user` VALUES ('12', 'huliang', '123456', '17542162634');
INSERT INTO `user` VALUES ('13', 'huliang', '123456', '17542162634');
INSERT INTO `user` VALUES ('15', 'huliang', '123456', '17542162634');
INSERT INTO `user` VALUES ('16', 'huliang', '123456', '17542162634');
INSERT INTO `user` VALUES ('17', 'huliang', '123456', '17542162634');
INSERT INTO `user` VALUES ('18', 'huliang', '123456', '17542162634');

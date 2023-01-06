/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : localhost:3306
 Source Schema         : pcr

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 05/01/2023 22:14:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for boss
-- ----------------------------
DROP TABLE IF EXISTS `boss`;
CREATE TABLE `boss`  (
  `bossid` int NOT NULL,
  `boss_value` int NOT NULL,
  `boss_syume` int NOT NULL,
  `stage_one` int NOT NULL,
  `stage_two` int NOT NULL,
  `stage_three` int NOT NULL,
  `stage_four` int NOT NULL,
  `stage_five` int NOT NULL,
  `now_max_value` int NOT NULL,
  `guashu` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `jinru` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`bossid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of boss
-- ----------------------------
INSERT INTO `boss` VALUES (1, 6000000, 1, 6000000, 6000000, 12000000, 22000000, 145000000, 6000000, ' ', ' ');
INSERT INTO `boss` VALUES (2, 8000000, 1, 8000000, 8000000, 14000000, 23000000, 150000000, 8000000, ' ', ' ');
INSERT INTO `boss` VALUES (3, 10000000, 1, 10000000, 10000000, 17000000, 27000000, 175000000, 10000000, ' ', ' ');
INSERT INTO `boss` VALUES (4, 12000000, 1, 12000000, 12000000, 19000000, 29000000, 195000000, 12000000, ' ', ' ');
INSERT INTO `boss` VALUES (5, 15000000, 1, 15000000, 15000000, 22000000, 31000000, 210000000, 15000000, ' ', ' ');

-- ----------------------------
-- Table structure for record
-- ----------------------------
DROP TABLE IF EXISTS `record`;
CREATE TABLE `record`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `target` int NOT NULL,
  `attackvalue` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 61 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of record
-- ----------------------------
INSERT INTO `record` VALUES (38, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (39, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (40, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (41, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (42, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (43, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (44, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (45, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (46, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (47, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (48, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (49, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (50, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (51, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (52, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (53, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (54, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (55, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (56, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (57, '示例数据', '2021-01-25 17:48:30', 1, 1);
INSERT INTO `record` VALUES (58, 'admin', '2023-01-05 18:55:55', 2, 900000);
INSERT INTO `record` VALUES (59, 'admin', '2023-01-05 18:56:05', 2, 7099700);
INSERT INTO `record` VALUES (60, '1234', '2023-01-05 19:27:18', 2, 3000000);

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `sessionid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `login` int NOT NULL,
  `expiration_time` datetime(6) NOT NULL,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 65 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sessions
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `authority` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (15, 'admin', 'admin', '管理员', 1);

SET FOREIGN_KEY_CHECKS = 1;

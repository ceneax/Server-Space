/*
 Navicat Premium Data Transfer

 Source Server         : mysql(local)
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : space

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 06/04/2020 12:50:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_comment`;
CREATE TABLE `t_comment`  (
  `comment_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `content` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '内容',
  `createdAt` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `updatedAt` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `replyedAt` datetime(0) NULL DEFAULT NULL COMMENT '回复日期',
  `ip` bigint(0) NULL DEFAULT NULL COMMENT 'IP地址',
  `haveReply` tinyint(0) NULL DEFAULT NULL COMMENT '是否有回复',
  `user_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '评论人ID',
  `likeCount` bigint(0) NULL DEFAULT NULL COMMENT '点赞计数',
  `hateCount` bigint(0) NULL DEFAULT NULL COMMENT '踩计数',
  PRIMARY KEY (`comment_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '评论' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_inform
-- ----------------------------
DROP TABLE IF EXISTS `t_inform`;
CREATE TABLE `t_inform`  (
  `inform_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `isHome` tinyint(1) NULL DEFAULT NULL COMMENT '是否首页提示',
  `type` int(0) NULL DEFAULT NULL COMMENT '公告类型',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '内容',
  `createdAt` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  PRIMARY KEY (`inform_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '公告' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_notice
-- ----------------------------
DROP TABLE IF EXISTS `t_notice`;
CREATE TABLE `t_notice`  (
  `notice_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '通知ID',
  `result_id` bigint(0) NULL DEFAULT NULL COMMENT '推送ID',
  `type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '通知类型，T被评论通知，R评论被回复通知，P回复被回复通知',
  `isRead` tinyint(1) NULL DEFAULT NULL COMMENT '是否已读',
  `user_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '推送用户ID',
  `createdAt` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `cause_id` bigint(0) NULL DEFAULT NULL COMMENT '引起通知的ID',
  `causeuser_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '引起通知的用户ID',
  PRIMARY KEY (`notice_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '通知' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_operation
-- ----------------------------
DROP TABLE IF EXISTS `t_operation`;
CREATE TABLE `t_operation`  (
  `oper_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '操作ID',
  `opername` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '操作名',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '编码',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '方法',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'URL',
  `pid` bigint(0) NULL DEFAULT NULL COMMENT '父操作ID',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '权限控制-操作' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_permission
-- ----------------------------
DROP TABLE IF EXISTS `t_permission`;
CREATE TABLE `t_permission`  (
  `per_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `pername` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '权限名',
  PRIMARY KEY (`per_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '权限控制-权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_permission_operation
-- ----------------------------
DROP TABLE IF EXISTS `t_permission_operation`;
CREATE TABLE `t_permission_operation`  (
  `per_id` bigint(0) NOT NULL COMMENT '权限ID',
  `oper_id` bigint(0) NOT NULL COMMENT '操作ID',
  PRIMARY KEY (`per_id`, `oper_id`) USING BTREE,
  INDEX `fk_po_oper_id`(`oper_id`) USING BTREE,
  CONSTRAINT `fk_po_oper_id` FOREIGN KEY (`oper_id`) REFERENCES `t_operation` (`oper_id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `fk_po_per_id` FOREIGN KEY (`per_id`) REFERENCES `t_permission` (`per_id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '权限控制-权限与操作关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_reply
-- ----------------------------
DROP TABLE IF EXISTS `t_reply`;
CREATE TABLE `t_reply`  (
  `reply_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '回复ID',
  `comment_id` bigint(0) NOT NULL COMMENT '评论ID',
  `content` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '内容',
  `createdAt` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `updatedAt` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `ip` bigint(0) NULL DEFAULT NULL COMMENT 'IP地址',
  `pid` bigint(0) NULL DEFAULT NULL COMMENT '父回复',
  `user_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '回复人员ID',
  `likeCount` bigint(0) NULL DEFAULT NULL COMMENT '点赞计数',
  `hateCount` bigint(0) NULL DEFAULT NULL COMMENT '踩计数',
  `deletedAt` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`reply_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '回复' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_rocket
-- ----------------------------
DROP TABLE IF EXISTS `t_rocket`;
CREATE TABLE `t_rocket`  (
  `rocket_id` bigint(0) NOT NULL COMMENT '火箭id',
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '图片',
  `chinesename` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '中文名',
  `englishname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '英文名',
  `type` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '类别',
  `RDAgency` bigint(0) NULL DEFAULT NULL COMMENT '研制单位',
  `LEO` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'LEO运载能力',
  `GTO` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'GTO运载能力',
  `height` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '箭体高度',
  `liftoffthrust` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '起飞推力',
  `liftoffquality` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '起飞质量',
  `status` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '状态',
  `isTop` tinyint(1) NULL DEFAULT NULL COMMENT '是否置顶',
  `commentCount` bigint(0) NULL DEFAULT NULL COMMENT '评论计数',
  `likeCount` bigint(0) NULL DEFAULT NULL COMMENT '点赞计数',
  `hateCount` bigint(0) NULL DEFAULT NULL COMMENT '踩计数',
  `shareCount` bigint(0) NULL DEFAULT NULL COMMENT '分享计数',
  `favoriteCount` bigint(0) NULL DEFAULT NULL COMMENT '收藏计数',
  `createdAt` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `updatedAt` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `commnetAt` datetime(0) NULL DEFAULT NULL COMMENT '最后评论日期',
  `deletedAt` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  `pid` bigint(0) NULL DEFAULT NULL COMMENT '基本型火箭ID',
  PRIMARY KEY (`rocket_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_role
-- ----------------------------
DROP TABLE IF EXISTS `t_role`;
CREATE TABLE `t_role`  (
  `role_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `rolename` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '角色名',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '权限控制-角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `t_role_permission`;
CREATE TABLE `t_role_permission`  (
  `role_id` bigint(0) NOT NULL COMMENT '角色ID',
  `per_id` bigint(0) NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`role_id`, `per_id`) USING BTREE,
  INDEX `fk_rp_per_id`(`per_id`) USING BTREE,
  CONSTRAINT `fk_rp_per_id` FOREIGN KEY (`per_id`) REFERENCES `t_permission` (`per_id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `fk_rp_role_id` FOREIGN KEY (`role_id`) REFERENCES `t_role` (`role_id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '权限控制-角色与权限关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_section
-- ----------------------------
DROP TABLE IF EXISTS `t_section`;
CREATE TABLE `t_section`  (
  `section_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '板块ID',
  `name` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '板块名称',
  `clickRate` bigint(0) NULL DEFAULT NULL COMMENT '点击率',
  `posts` bigint(0) NULL DEFAULT NULL COMMENT '发帖数',
  `delflag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标记',
  `classfication` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '分类',
  PRIMARY KEY (`section_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '版块' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_topic
-- ----------------------------
DROP TABLE IF EXISTS `t_topic`;
CREATE TABLE `t_topic`  (
  `topic_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '主题ID',
  `section_id` bigint(0) NOT NULL COMMENT '板块ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '标题',
  `shortContent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '短内容',
  `user_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户ID',
  `isTop` tinyint(1) NULL DEFAULT NULL COMMENT '是否置顶',
  `commentCount` bigint(0) NULL DEFAULT NULL COMMENT '评论计数',
  `likeCount` bigint(0) NULL DEFAULT NULL COMMENT '点赞计数',
  `hateCount` bigint(0) NULL DEFAULT NULL COMMENT '踩计数',
  `shareCount` bigint(0) NULL DEFAULT NULL COMMENT '分享计数',
  `favoriteCount` bigint(0) NULL DEFAULT NULL COMMENT '收藏计数',
  `createdAt` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `updatedAt` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `commnetAt` datetime(0) NULL DEFAULT NULL COMMENT '最后评论日期',
  `deletedAt` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`topic_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '主题' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_topic_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_topic_comment`;
CREATE TABLE `t_topic_comment`  (
  `topic_id` bigint(0) NOT NULL COMMENT '主题ID',
  `comment_id` bigint(0) NOT NULL COMMENT '评论ID',
  `deletedAt` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`topic_id`, `comment_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '主题与评论关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_topic_content
-- ----------------------------
DROP TABLE IF EXISTS `t_topic_content`;
CREATE TABLE `t_topic_content`  (
  `topic_id` bigint(0) NOT NULL COMMENT '主题ID',
  `section_id` bigint(0) NOT NULL COMMENT '板块ID',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '内容',
  PRIMARY KEY (`topic_id`, `section_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '主题内容' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
  `user_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户ID',
  `username` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户名',
  `password` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '密码',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '头像',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户基本信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user_evaluate
-- ----------------------------
DROP TABLE IF EXISTS `t_user_evaluate`;
CREATE TABLE `t_user_evaluate`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `user_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户ID',
  `section_id` bigint(0) NULL DEFAULT NULL COMMENT '板块ID',
  `type_id` bigint(0) NULL DEFAULT NULL COMMENT '主题或评论的ID',
  `type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '类型，1主题点赞，2，评论点赞，3，回复点赞',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态，-1取消评价，1点赞，点踩',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user_favorite
-- ----------------------------
DROP TABLE IF EXISTS `t_user_favorite`;
CREATE TABLE `t_user_favorite`  (
  `id` bigint(0) NOT NULL COMMENT '主键ID',
  `user_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户ID',
  `section_id` bigint(0) NULL DEFAULT NULL COMMENT '板块ID',
  `type_id` bigint(0) NULL DEFAULT NULL COMMENT '收藏对应的ID',
  `type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '收藏类型，1-主题收藏',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '收藏状态，-1-取消收藏，1-收藏',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user_info
-- ----------------------------
DROP TABLE IF EXISTS `t_user_info`;
CREATE TABLE `t_user_info`  (
  `user_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户ID',
  `sex` tinyint(1) NULL DEFAULT NULL COMMENT '性别，1男，2女',
  `emailVerified` tinyint(1) NULL DEFAULT NULL COMMENT '是否验证邮箱',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '邮箱',
  `createdAt` datetime(0) NULL DEFAULT NULL COMMENT '创建于',
  `updatedAt` datetime(0) NULL DEFAULT NULL COMMENT '更新于',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户详细信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user_role
-- ----------------------------
DROP TABLE IF EXISTS `t_user_role`;
CREATE TABLE `t_user_role`  (
  `user_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户ID',
  `role_id` bigint(0) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '权限控制-用户与角色关系' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

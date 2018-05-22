/*
Navicat MySQL Data Transfer

Source Server         : root_sub
Source Server Version : 50505
Source Host           : 192.168.0.25:3306
Source Database       : rbac

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2018-05-22 16:27:24
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for access
-- ----------------------------
DROP TABLE IF EXISTS `access`;
CREATE TABLE `access` (
  `role_id` smallint(6) NOT NULL COMMENT '角色ID',
  `permission_id` smallint(6) NOT NULL COMMENT '权限节点ID',
  `state` tinyint(1) NOT NULL COMMENT '状态位',
  `is_delete` tinyint(1) NOT NULL COMMENT '软删除标识',
  `deleted_at` int(11) NOT NULL COMMENT '软删除时间',
  `created_at` int(11) NOT NULL COMMENT '创建时间',
  `updated_at` int(11) NOT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='@角色与权限节点(授权数据表)';

-- ----------------------------
-- Records of access
-- ----------------------------

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `permission_id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '权限节点ID',
  `parent_id` smallint(6) NOT NULL COMMENT '父级权限节点ID',
  `name` varchar(20) NOT NULL COMMENT '权限节点名称',
  `path` varchar(100) NOT NULL COMMENT '权限节点路径',
  `remark` varchar(255) NOT NULL COMMENT '备注',
  `level` tinyint(1) unsigned NOT NULL COMMENT '权限节点级别',
  `sort` smallint(6) unsigned NOT NULL COMMENT '排序号',
  `state` tinyint(1) unsigned NOT NULL COMMENT '状态',
  `is_delete` tinyint(1) unsigned NOT NULL COMMENT '软删除标识',
  `deleted_at` int(11) NOT NULL COMMENT '软删除时间',
  `created_at` int(11) NOT NULL COMMENT '创建时间',
  `updated_at` int(11) NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='@权限节点数据表';

-- ----------------------------
-- Records of permission
-- ----------------------------

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `role_id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `parent_id` smallint(6) NOT NULL COMMENT '父级角色ID',
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `remark` varchar(100) NOT NULL COMMENT '备注',
  `state` tinyint(1) NOT NULL COMMENT '状态',
  `is_delete` tinyint(1) NOT NULL COMMENT '软删除标识',
  `deleted_at` int(11) NOT NULL COMMENT '软删除标识',
  `created_at` int(11) NOT NULL COMMENT '创建时间',
  `updated_at` int(11) NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='@角色数据表';

-- ----------------------------
-- Records of role
-- ----------------------------

-- ----------------------------
-- Table structure for role_user
-- ----------------------------
DROP TABLE IF EXISTS `role_user`;
CREATE TABLE `role_user` (
  `role_id` smallint(6) NOT NULL COMMENT '角色ID',
  `user_id` smallint(6) NOT NULL COMMENT '用户ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='@角色与用户(关联数据表)';

-- ----------------------------
-- Records of role_user
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `account` char(32) NOT NULL COMMENT '账号',
  `realname` varchar(255) NOT NULL COMMENT '真实姓名',
  `password` char(32) NOT NULL COMMENT '密码',
  `last_login_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_login_ip` char(15) NOT NULL COMMENT '最后登录IP',
  `login_count` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '登录账号',
  `email` varchar(50) NOT NULL COMMENT '邮箱',
  `mobile` char(11) NOT NULL COMMENT '手机号码',
  `remark` varchar(255) NOT NULL COMMENT '备注',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态位',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '软删除标识',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL COMMENT '软删除时间',
  PRIMARY KEY (`user_id`),
  KEY `accountpassword` (`account`,`password`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='@用户数据表';

-- ----------------------------
-- Records of user
-- ----------------------------

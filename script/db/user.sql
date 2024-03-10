USE `gomyrss`;

-- 用户表
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,
    `username`   VARCHAR(50) UNIQUE NOT NULL COMMENT '用户名, 唯一',
    `password`   VARCHAR(50)        NOT NULL COMMENT '密码, 明文保存',
    `email`      VARCHAR(50) COMMENT '邮箱, 明文保存',
    `is_admin`   TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否为管理员',
    `created_at` DATETIME           NOT NULL COMMENT '创建时间',
    `updated_at` DATETIME           NOT NULL COMMENT '更新时间',
    `delete_at`  DATETIME           NOT NULL COMMENT '删除时间'
);

-- 权限表
DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,
    `name`       VARCHAR(50) NOT NULL COMMENT '权限名',
    `created_at` DATETIME    NOT NULL COMMENT '创建时间',
    `updated_at` DATETIME    NOT NULL COMMENT '更新时间',
    `delete_at`  DATETIME    NOT NULL COMMENT '删除时间'
);

-- 用户权限关联表
DROP TABLE IF EXISTS `user_permission_rel`;
CREATE TABLE `user_permission_rel`
(
    `user_id`       INT      NOT NULL COMMENT '用户id',
    `permission_id` INT      NOT NULL COMMENT '权限id',
    `created_at`    DATETIME NOT NULL COMMENT '创建时间',
    `updated_at`    DATETIME NOT NULL COMMENT '更新时间',
    `delete_at`     DATETIME NOT NULL COMMENT '删除时间'
);

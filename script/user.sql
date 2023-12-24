USE playground;

# ACL 权限控制，admin 用户默认分配所有权限，其他用户的权限由 admin 分配
# 分配权限给其他用户时不做其他的限制，相当于 admin 是一个拥有所有权限的普通用户

-- 用户表
CREATE TABLE `users`
(
    `id`       INT PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(50) UNIQUE NOT NULL COMMENT '用户名, 唯一',
    `password` VARCHAR(50)        NOT NULL COMMENT '密码, 明文保存',
    `email`    VARCHAR(50)                 DEFAULT NULL COMMENT '邮箱, 明文保存',
    `is_admin` TINYINT(1)         NOT NULL DEFAULT 0 COMMENT '是否为管理员'
);

-- 权限表
CREATE TABLE `permissions`
(
    `id`   INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(50) NOT NULL COMMENT '权限名'
);

-- 用户权限关联表
CREATE TABLE `user_permission_rel`
(
    `user_id`       INT NOT NULL COMMENT '用户Id',
    `permission_id` INT NOT NULL COMMENT '权限Id',
    PRIMARY KEY (`user_id`, `permission_id`)
);

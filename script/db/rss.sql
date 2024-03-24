USE `gomyrss`;

-- 订阅表
DROP TABLE IF EXISTS `rss`;
CREATE TABLE `rss`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键',
    `title`       VARCHAR(50)   NOT NULL COMMENT '标题',
    `feed_rul`    VARCHAR(200)  NOT NULL COMMENT '订阅 url',
    `website_url` VARCHAR(200)  NOT NULL COMMENT '网站 url',
    `description` VARCHAR(1000) NOT NULL COMMENT '描述',
    `type`        INT           NOT NULL COMMENT '订阅类型',
    `created_at`  DATETIME      NOT NULL COMMENT '创建时间',
    `updated_at`  DATETIME      NOT NULL COMMENT '更新时间',
    `delete_at`   DATETIME      NOT NULL COMMENT '删除时间'
);

-- 订阅分组表
DROP TABLE IF EXISTS `feed_group`;
CREATE TABLE `feed_group`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键',
    `group_name` VARCHAR(50) NOT NULL COMMENT '订阅分组名称',
    `created_at` DATETIME    NOT NULL COMMENT '创建时间',
    `updated_at` DATETIME    NOT NULL COMMENT '更新时间',
    `delete_at`  DATETIME    NOT NULL COMMENT '删除时间'
);

-- 订阅分组关联表
DROP TABLE IF EXISTS `rss_feed_group_rel`;
CREATE TABLE `rss_feed_group_rel`
(
    `rss_id`        INT      NOT NULL COMMENT '订阅 id',
    `feed_group_id` INT      NOT NULL COMMENT '订阅分组 id',
    `created_at`    DATETIME NOT NULL COMMENT '创建时间',
    `updated_at`    DATETIME NOT NULL COMMENT '更新时间',
    `delete_at`     DATETIME NOT NULL COMMENT '删除时间'
);

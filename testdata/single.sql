CREATE TABLE `user`
(
    `id`              int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username`        varchar(200)     NOT NULL COMMENT 'username',
    `first_name`      varchar(100)              DEFAULT '' COMMENT 'first name',
    `last_name`       varchar(100)              DEFAULT '',
    `password`        varchar(512)     NOT NULL COMMENT '密码',
    `email`           varchar(100)     NOT NULL DEFAULT '' COMMENT 'EMAIL',
    `status`          tinyint(1)                DEFAULT '1' COMMENT 'status',
    `last_login_time` int(11)          NOT NULL DEFAULT '0' COMMENT '最后登录时间',
    `last_login_ip`   varchar(16)               DEFAULT NULL COMMENT '最后登录IP',
    `update_time`     int(11)          NOT NULL DEFAULT '0' COMMENT '更新时间',
    `create_time`     int(11)          NOT NULL DEFAULT '0' COMMENT '创建时间',
    `delete_time`     int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='系统用户表'
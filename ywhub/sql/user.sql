CREATE TABLE `user`
(
    `id`           int(10) unsigned                                              NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`      int(10) unsigned                                                       DEFAULT NULL COMMENT '用户id',
    `user_name`    varchar(20)   NOT NULL COMMENT '用户名',
    `pwd`          varchar(128)   NOT NULL COMMENT '密码',
    `email`          varchar(50)   NOT NULL COMMENT 'EMail',
    `created_time` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` timestamp ,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_user_name_uindex` (`user_name`),
    UNIQUE KEY `user_user_id_uindex` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='用户表';
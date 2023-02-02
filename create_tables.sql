Create Database If Not Exists redrock;
create table users(
                      id bigint not null auto_increment comment 'id,用于点赞接口',
                      uid bigint not null comment '用户id',
                      username varchar(50) not null comment '用户名',
                      password varchar(50) not null comment '密码',
                      email varchar(100) not null comment '邮箱',
                      create_time timestamp null default current_timestamp,
                      update_time timestamp null default current_timestamp on update current_timestamp,
                      gender smallint NOT NULL DEFAULT 0 COMMENT '性别，未设定时为0，男性为1，女性为2',
                      introduction varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '简介',
                      `head_portrait` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '头像',
                      `background_img` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  DEFAULT '' COMMENT '背景',
                      primary key (id),
                      unique key `idx_uid` (uid) using btree ,
                      unique key `idx_username` (username) using btree
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;
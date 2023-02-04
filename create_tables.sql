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

create table comments(
    cid bigint not null comment '评论id',
    author_id bigint not null comment '评论人id',
    book_id bigint not null comment '被评书籍id',
    parent_id bigint not null default 1 comment '父级评论id,若为最高级评论则为1',
    root_id bigint not null default 1 comment '根级评论id，若为最高级评论则为1    ',
    commented_uid bigint not null comment '被回复的人的id',
    stars bigint not null default 0 comment '点赞数',
    content varchar(2056) not null comment '内容,如果评论被删除则此字段为已删除',
    create_time timestamp null default current_timestamp,
    primary key (cid),
    key `idx_author_id` (author_id) using btree ,
    key `idx_book_id` (book_id) using btree
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;

create table tags (
    tid smallint not null comment '标签id',
    tag varchar(128) not null comment '标签名',
    primary key (tid)
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;

INSERT INTO `tags` VALUES (1, '玄乎');
INSERT INTO `tags` VALUES (2, '奇幻');
INSERT INTO `tags` VALUES (3, '武侠');
INSERT INTO `tags` VALUES (4, '仙侠');
INSERT INTO `tags` VALUES (5, '都市');
INSERT INTO `tags` VALUES (6, '现实');
INSERT INTO `tags` VALUES (7, '军事');
INSERT INTO `tags` VALUES (8, '历史');
INSERT INTO `tags` VALUES (9, '游戏');
INSERT INTO `tags` VALUES (10, '体育');
INSERT INTO `tags` VALUES (11, '科幻');
INSERT INTO `tags` VALUES (12, '诸天无限');
INSERT INTO `tags` VALUES (13, '悬疑');
INSERT INTO `tags` VALUES (14, '轻小说');
INSERT INTO `tags` VALUES (15, '短篇');

create table book_tag_map(
    id bigint not null  auto_increment ,
    tid smallint not null comment '标签id',
    bid bigint not null comment '书籍id',
    primary key (id),
    key `idx_bid` (bid) using btree ,
    key `idx_tid` (tid) using btree
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;

INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (5, 2);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (4, 3);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (1, 4);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (14, 5);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (14, 6);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (4, 7);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (11, 8);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (5, 9);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (8, 10);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (13, 11);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (11, 12);
INSERT INTO `redrock`.`book_tag_map` (`tid`, `bid`) VALUES (11, 1);


create table books(
    bid bigint not null auto_increment comment '书籍id',
    name varchar(128) not null comment '书名',
    author varchar(128) not null comment '作者',
    score bigint default 0 not null comment '得分',
    cover varchar(1024) not null comment '封面url',
    link varchar(1024) not null comment '书籍连接',
    publish_time timestamp null default current_timestamp comment '发布时间',
    author_id smallint default 0 not null ,
    primary key (bid),
    key `idx_name` (name) using btree
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;

INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (1, '灵境行者', '卖报小郎君', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1031940621/180', 'https://book.qidian.com/info/1031940621/', '2023-02-04 15:37:23');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (2, '深空彼岸', '辰东', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1027339371/180', 'https://book.qidian.com/info/1027339371/', '2023-02-04 15:38:24');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (3, '光阴之外', '耳根', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1031777108/180', 'https://book.qidian.com/info/1031777108/', '2023-02-04 07:38:51');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (4, '道诡异仙', '狐尾的笔', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1031794030/180', 'https://book.qidian.com/info/1031794030/', '2023-02-04 07:39:15');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (5, '轮回乐园', '那一只蚊子', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1009817672/180', 'https://book.qidian.com/info/1009817672/', '2023-02-04 07:39:41');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (6, '这游戏也太真实了', '晨星LL', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1029391348/180', 'https://book.qidian.com/info/1029391348/', '2023-02-04 07:40:01');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (7, '择日飞升', '宅猪', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1032778366/180', 'https://book.qidian.com/info/1032778366/', '2023-02-04 07:40:33');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (8, '深海余烬', '远瞳', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1034360760/180', 'https://book.qidian.com/info/1034360760/', '2023-02-04 07:40:59');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (9, '明克街13号', '纯洁滴小龙', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1030870265/180', 'https://book.qidian.com/info/1030870265/', '2023-02-04 07:41:32');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (10, '唐人的餐桌', '孑与2', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1033914374/180', 'https://book.qidian.com/info/1033914374/', '2023-02-04 07:41:56');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (11, '我的治愈系游戏', '我会修空调', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1025901449/180', 'https://book.qidian.com/info/1025901449/', '2023-02-04 07:42:34');
INSERT INTO `redrock`.`books` (`bid`, `name`, `author`, `score`, `cover`, `link`, `publish_time`) VALUES (12, '\r\n深渊独行', '言归正传', 0, 'https://bookcover.yuewen.com/qdbimg/349573/1034695389/180', 'https://book.qidian.com/info/1034695389/', '2023-02-04 07:43:17');


create table bookshelf(
    id bigint not null auto_increment,
    uid bigint not null comment '用户id',
    bid bigint not null  comment '书籍id',
    primary key (id),
    key `idx_uid` (uid) using btree,
    key `idx_bid` (bid) using btree
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;
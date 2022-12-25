create database if not exists Books;

create table users(
                      id bigint not null auto_increment,
                      username varchar(30) collate utf8mb4_general_ci not null,
                      password varchar(50) collate utf8mb4_general_ci not null,
                      createAt timestamp null default current_timestamp,
                      updateAt timestamp null default current_timestamp on update current_timestamp,
                      primary key (id),
                      unique key `idx_username` (username) using btree
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;

create table books (
    id bigint not null auto_increment,
    name varchar(30) collate utf8mb4_general_ci not null,
    author varchar(30) collate utf8mb4_general_ci not null,
    status smallint default 0, #0为未借阅，1为已借阅
    overdue smallint default 0, #0为未逾期，1为逾期
    borrower_id bigint,
    borrowAt timestamp ,
    primary key (id),
    unique key `idx_borrower_id` (borrower_id) using btree
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;
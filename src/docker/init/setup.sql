create database gochat;
\connect gochat;

drop table if exists posts;

create table posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);

insert into posts (content, author) values ('test_content', 'test_author')
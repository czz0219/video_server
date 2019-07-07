create database video_server;
use video_server;
create table users(
	id  int(10) unsigned primary key auto_increment,
	login_name varchar(64) unique key,
	pwd text not null
);
create table video_info(
id varchar(64) primary key not null,
author_id int(10) unsigned,
name text,
display_ctime text,
create_time datetime
);
create table comments(
id varchar(64) primary key not null,
video_id varchar(64),
author_id int(10) unsigned,
content text,
time datetime default current_timestamp
);
create table sessions(
session_id varchar(128) primary key not null,
ttl tinytext,
login_name text
);
create table video_del_rec(
	video_id varchar(64) primary key not null
);
/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2022/5/15 9:59:14                            */
/*==============================================================*/


drop table if exists dy_comment;

drop table if exists dy_favorite;

drop table if exists dy_feed;

drop table if exists dy_follow;

drop table if exists dy_user;

drop table if exists dy_video;

/*==============================================================*/
/* Table: dy_comment                                            */
/*==============================================================*/
create table dy_comment
(
   id                   int not null auto_increment,
   sender_id            int not null,
   video_id             int not null,
   content              text not null,
   status               int not null,
   create_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) not null,
   update_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) not null,
   primary key (id)
)auto_increment=1;

ALTER TABLE dy_comment ADD INDEX user_id (sender_id);
ALTER TABLE dy_comment ADD INDEX video_id (video_id);


/*==============================================================*/
/* Table: dy_favorite                                           */
/*==============================================================*/
create table dy_favorite
(
   video_id             int not null,
   user_id              int not null,
   create_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) not null,
   update_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) not null,
   primary key (video_id, user_id)
);

/*==============================================================*/
/* Table: dy_feed                                               */
/*==============================================================*/
create table dy_feed
(
   user_id              int not null,
   video_id             int not null,
   create_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) not null,
   update_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) not null,
   primary key (user_id, video_id)
);

/*==============================================================*/
/* Table: dy_follow                                             */
/*==============================================================*/
create table dy_follow
(
   follow_id             int not null,
   follower_id           int not null,
   create_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) not null,
   update_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) not null,
   primary key (follow_id, follower_id)
);

/*==============================================================*/
/* Table: dy_user                                               */
/*==============================================================*/
create table dy_user
(
   id                   int not null auto_increment,
   name                 varchar(20) not null unique,
   follow_count         int not null,
   follower_count       int not null,
   password             varchar(32) not null,
   create_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) not null,
   update_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) not null,
   primary key (id),
)auto_increment=1;

/*==============================================================*/
/* Table: dy_video                                              */
/*==============================================================*/ 
create table dy_video
(
   id                   int not null auto_increment,
   title				varchar(50) not null,
   play_url             varchar(100) not null,
   cover_url            varchar(100) not null,
   favorite_count       int not null,
   comment_count        int not null,
   author_id            int not null,
   create_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) not null,
   update_time TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) not null,
   primary key (id)
)auto_increment=1;

ALTER TABLE dy_video ADD INDEX user_id (author_id);


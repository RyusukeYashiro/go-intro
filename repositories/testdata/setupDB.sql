-- #記事データを格納するためのテーブル
create table if not exists articles (
    article_id integer unsigned auto_increment primary key,
    title varchar(255) not null,
    contents text not null,
    username varchar(255) not null,
    nice integer not null,
    created_at datetime DEFAULT CURRENT_TIMESTAMP
);

-- コメントデータを格納するためのテーブル
create table if not exists comments (
    comment_id integer unsigned auto_increment primary key,
    article_id integer unsigned not null,
    message text not null,
    created_at datetime DEFAULT CURRENT_TIMESTAMP
);


insert into articles (title , contents , username , nice , created_at) values 
	("firstPost" , "this is my first blog" , "yashiro" , 2 , now());

insert into articles (title, contents, username, nice) values
('2nd', 'Second blog post', 'saki', 4);

insert into comments (article_id, message, created_at) values
(1, '1st comment yeah', now());

insert into comments (article_id, message) values
(1, 'welcome');
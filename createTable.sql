#記事データを格納するためのテーブル
create table if not exists articles (
	article_id integer unsigned auto_increment primary key,
	title varchar(255) not null,
	contents text not null,
	username varchar(255) not null,
	nice integer not null,
	created_at datetime
);

#コメントデータを格納するためのテーブル
create table if not exists comments (
	comment_id integer unsigned auto_increment primary key,
	article_id integer unsigned not null,
	message text not null,
	created_at datetime
)

insert into articles (title , contents , username , nice , created_at) values 
	("firstPost" , "this is my first blog" , "yashiro" , 2 , now());

insert into articles (title, contents, username, nice) values
('2nd', 'Second blog post', 'saki', 4);

insert into comments (article_id, message, created_at) values
(1, '1st comment yeah', now());

insert into comments (article_id, message) values
(1, 'welcome');
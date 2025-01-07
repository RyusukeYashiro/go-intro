package models

import "time"

// ブログに対してのコメント構造体
// jsonエンコード時にjsonキーを指定する<-jsonはスネークケースが基本だから
type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice_num"`
	CommentList []Comment `json:"comment_list"`
	CreatedAt   time.Time `json:"created_at"`
}

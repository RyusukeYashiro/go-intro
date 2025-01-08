// モックデータを使う
package models

import "time"

var (
Comment1 = Comment{
	CommentID: 1,
	ArticleID: 1,
	Message:   "This is the first comment.",
	CreatedAt: time.Now(),
}
Comment2 = Comment{
	CommentID: 2,
	ArticleID: 1,
	Message:   "Another perspective on the topic.",
	CreatedAt: time.Now(),
}
Comment3 = Comment{
	CommentID: 3,
	ArticleID: 1,
	Message:   "Great article! Thanks for sharing.",
	CreatedAt: time.Now(),
}
Comment4 = Comment{
	CommentID: 4,
	ArticleID: 2,
	Message:   "Interesting insights. Keep it up!",
	CreatedAt: time.Now(),
}
Comment5 = Comment{
	CommentID: 5,
	ArticleID: 2,
	Message:   "I have a different opinion on this.",
	CreatedAt: time.Now(),
}
)

var (
	Article1 = Article{
		ID:         1,
		Title:      "Understanding Go Structs",
		Contents:   "In this article, we delve deep into the concept of structs in Go.",
		UserName:   "john_doe",
		NiceNum:    7,
		CommentList: []Comment{Comment1, Comment2, Comment3},
		CreatedAt:  time.Now(),
	}
	
	Article2 = Article{
		ID:         2,
		Title:      "Concurrency in Go",
		Contents:   "Go makes concurrency easy with goroutines and channels.",
		UserName:   "jane_smith",
		NiceNum:    4,
		CommentList: []Comment{Comment4, Comment5},
		CreatedAt:  time.Now(),
	}
)

var (AllComments = []Comment{
	Comment1,Comment2,Comment3,Comment4,Comment5,
})
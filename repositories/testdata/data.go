package testdata

import "go-intro/models"

var ArticleTestData = []models.Article{
	models.Article{
	ID: 1,
	Title: "firstPost",
	Contents: "this is my first blog",
	UserName: "yashiro",
	NiceNum: 3,
	},
	models.Article{
	ID: 2,
	Title: "2nd",
	Contents: "Second blog post",
	UserName: "saki",
	NiceNum: 4,
	},
	}
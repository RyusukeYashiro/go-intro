package services

import (
	"go-intro/models"
	"go-intro/repositories"
	"log"
)

// ArticleDetailHandler 内: 指定 ID の記事をデータベースから取得する
func GetArticleService(articleID int) (models.Article , error) {
	db , err := connectDB()
	if err != nil {
		return models.Article{},  err
	}
	defer db.Close()	
	// 1. repositories 層の関数 SelectArticleDetail で記事の詳細を取得
	article , err := repositories.SelectArticleDetail(db , articleID)
	if err != nil {
		return models.Article{} , err
	}
	// 2. repositories 層の関数 SelectCommentList でコメント一覧を取得
	commentList , err := repositories.SelectCommentList(db , articleID)
	if err != nil {
		return models.Article{} , err
	}
	article.CommentList = append(article.CommentList, commentList...)
	return article , nil
}

// 1 の内容をデータベースに挿入して、実際にデータベース内に収められた値を得る
func PostArticleService(article models.Article) (models.Article , error) {
	db , err := connectDB()
	if err != nil {
		return models.Article{} , err
	}
	defer db.Close()
	newArticle , err := repositories.InsertArticle(db , article)
	if err != nil {
		return models.Article{} , err
	}
	return newArticle , err	
}

// . クエリパラメータで指定されたページの記事一覧をデータベースから取得する
func GetArticleListService(page int) ([]models.Article , error) {
	db , err := connectDB()
	if err != nil {
		log.Printf("DB connection error: %v", err)  // 接続エラーのログ
		return nil , err
	}
	defer db.Close()
	articleList , err := repositories.SelectArticleList(db , page)
	if err != nil {
		log.Printf("Select article error: %v", err)  // クエリエラーのログ
		return nil , err
	}
	return articleList , err
}

func PostNiceService(article models.Article) (models.Article , error) {
	db , err := connectDB()
	if err != nil {
		return models.Article{} , err
	}
	defer db.Close()
	err = repositories.UpdateNiceNum(db , article.ID)
	if err != nil {
		return models.Article{} , err
	}

	article ,err = repositories.SelectArticleDetail(db ,  article.ID)
	if err != nil {
		return models.Article{} , err
	}
	return article , nil
}
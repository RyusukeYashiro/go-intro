// 新規投稿をデータベースにinsertする関数
package repositories

import (
	"database/sql"
	"fmt"
	"go-intro/models"
	"log"
)

// • POST /article: リクエストボディで受け取った記事を投稿する
func InsertArticle(db *sql.DB ,  article models.Article) (models.Article , error) {
	
	var newArticle models.Article
	const sqlStr = `insert into articles (title , contents , username , nice , created_at) values (? , ? , ? , 0  , now());`
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName
	result , err := db.Exec(sqlStr , article.Title , article.Contents , article.UserName)
	if err != nil {
		return models.Article{} , err
	}
	id , err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	newArticle.ID = int(id)
	return newArticle , nil
}

// GET /article/list: クエリパラメータ page で指定されたページ (1 ページに 5 個の記事
	// を表示) に表示するための記事一覧を取得する
func SelectArticleList(db *sql.DB , page int) ([]models.Article , error) {
	const sqlstr = `select article_id , title , contents , username , nice from articles limit ? offset ?`
	rows , err := db.Query(sqlstr , 5 , (page - 1) * 5); if err != nil {
		return nil , err
	}
	defer rows.Close()

	arrayArticles := make([]models.Article , 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID , &article.Title , &article.Contents , &article.UserName , &article.NiceNum)
		arrayArticles = append(arrayArticles, article)
	}
	return arrayArticles , nil
}

func SelectArticleDetail(db *sql.DB , articleID int) (models.Article , error) {
	const sqlStr = `select * from article where id = ?`
	row := db.QueryRow(sqlStr , articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return models.Article{} , err
	}
	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID , &article.Title , &article.Contents , &article.NiceNum , &article.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return models.Article{} , err
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	return article , nil
}

func addNiceNum(db *sql.DB , articleID int) (error) {
	// トランザクション処理
	tx , err := db.Begin(); if err != nil {
		fmt.Println(err)
		return  err
	}
	//指定されたidの記事を取得
	const sqlGetNice = `select nice from articles where article = ?`
	row := tx.QueryRow(sqlGetNice , articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return  err
	}

	var niceNum int
	err = row.Scan(&niceNum); if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice =  ? where article_id`
	_ , err = tx.Exec(sqlUpdateNice , niceNum + 1 , articleID)
	if err != nil {
		tx.Rollback()
		return  err
	}
	tx.Commit()
	return err
}
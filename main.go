// main.go
package main

import (
	"database/sql"
	"fmt"
	"go-intro/handlers"
	"go-intro/models"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//1. 第 2 引数 req *http.Request の中身を使って、レスポンスの中身を作成する
// 2. 作成したレスポンスの中身を、第一引数 w http.ResponseWriter に書き込む


func main() {
	// サーバーが受けとったhttpリクエストをどのハンドラに処理をさせるか決めるルーター
	r := mux.NewRouter()

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// open関数でdbに接続。その際、ドライバーを設定
	db , err := sql.Open("mysql" , dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
// sql.DB 型の Ping メソッドで疎通確認
	// if err := db.Ping(); err != nil {
	// 	fmt.Println(err) 
	// } else {
	// 	fmt.Println("connect! to DB")
	// }

	//データを挿入する処理
	article1 := models.Article{
		Title:    "insert test",
		Contents: "Can I insert data?",
		UserName: "test-user",
	}
	const sqlstr = `insert into articles (title , contents , username , nice , created_at) values(? , ? , ? , 0  , now());`
	result , err := db.Exec(sqlstr , article1.Title , article1.Contents , article1.UserName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	articleID := 1
	const sqlStr = `select * from articles where article_id = ?;`
	//最大でも１つの列を返すようなクエリを実行するときに使う=db.QueryRow
	row := db.QueryRow(sqlStr , articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// 変数 article の各フィールドに、取得レコードのデータを入れる
	var article models.Article
	var createdTime sql.NullTime
	// rows の中に格納されている取得レコード内容を読み出す
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n" , article)
	// net/httpパッケージ内で定義されているHandleFuncを用いる。パンドラ登録作業
	// http.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticle)
	// gorilla/mux では、受け付けていないメソッドのリクエストが来た場合には、ハン
	// ドラに処理を回す前にルータ自身で 405 エラーを返してくれる
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment" , handlers.PostComment).Methods(http.MethodPost)
	log.Println("server start at port 8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// ルータを使うときの処理。ここではListen関数の第２引数はルータの指定
	log.Fatal(http.ListenAndServe(":8080", r))
}

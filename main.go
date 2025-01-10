// main.go
package main

import (
	"database/sql"
	"fmt"
	"go-intro/handlers"
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

// main.go
package main

import (
	"go-intro/handlers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

//1. 第 2 引数 req *http.Request の中身を使って、レスポンスの中身を作成する
// 2. 作成したレスポンスの中身を、第一引数 w http.ResponseWriter に書き込む


func main() {
	// サーバーが受けとったhttpリクエストをどのハンドラに処理をさせるか決めるルーター
	r := mux.NewRouter()

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

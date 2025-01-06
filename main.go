// main.go
package main

import (
	"io"
	"log"
	"net/http"
)

//1. 第 2 引数 req *http.Request の中身を使って、レスポンスの中身を作成する
// 2. 作成したレスポンスの中身を、第一引数 w http.ResponseWriter に書き込む

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// ハンドラの第一引数として渡されていた http.ResponseWriter 型の変数 w に
		// "Hello, World!"と書き込む
		io.WriteString(w, "Hello world!\n")
	}
	// net/httpパッケージ内で定義されているHandleFuncを用いる。パンドラ登録作業
	http.HandleFunc("/", helloHandler)
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

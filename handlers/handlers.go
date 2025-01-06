// Go では main 以外のパッケージ名は、基本的にそのファイルが格納されているディレクトリ名と
// 同名にする必
package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

// Go では、他のパッケージからも参照可能な関数・変数・定数を作成するためには、その名前を大
// 文字から始める
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello gowold!!")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// ブログ記事の投稿をするためのエンドポイント
func PostArticle(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// クエリパラメーター取得機能
	queryMap := req.URL.Query()

	var pg int
	//ここのokでmap内のkeyに正しく値が入っているかどうかを確認
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		pg, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query param", http.StatusBadRequest)
			return
		}
	} else {
		pg = 1
	}

	reString := fmt.Sprintf(("Article List (pg %d)\n"), pg)
	io.WriteString(w, reString)

}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// artcile_Id := 1
	// パラメータの取得
	// Vars関数は引数にhttpリクエスト構造体をとった上でそこに含まれるパラメータをマップで返す
	// • mux.Vars(req) と指定した場合に得られる値: map[id:1]
	// mux.Vars(req)["id"] と指定した場合に得られる値: "1"
	artcile_Id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", artcile_Id)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostComment(w http.ResponseWriter, trq *http.Request) {
	io.WriteString(w, "Posting comments...\n")
}

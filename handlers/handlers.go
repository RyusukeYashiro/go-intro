// Go では main 以外のパッケージ名は、基本的にそのファイルが格納されているディレクトリ名と
// 同名にする必
package handlers

import (
	"encoding/json"
	"fmt"
	"go-intro/models"
	"io"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
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
	// io.WriteString(w, "Posting Nice...\n")
	// データ取得→ json エンコード」	
	article := models.Article1
	jsonData , err := json.MarshalIndent(article , "" , " ")
	if err != nil {
		http.Error(w , "fail to encode json\n" , http.StatusInternalServerError)
		return
	}
	// http.ResponseWriterには []byte を引数にとって書き込み処理を行うことができる Write
	// メソッドが備わっている
	w.Write(jsonData)	
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

	articleList := []models.Article{models.Article1 , models.Article2}
	jsonData , err := json.MarshalIndent(articleList , "" , " ")
	if err != nil {
		errMs := fmt.Sprintf(("Article List (pg %d)\n"), pg)
		http.Error(w , errMs , http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
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
	article := models.Article1
	jsonData , err := json.MarshalIndent(article , "" , " ")
	if err != nil {
		errMs := fmt.Sprintf("faild to encode json (articleID :%d)\n" , artcile_Id)
		http.Error(w ,errMs  , http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w , "Posting nice...\n")
	var articleWithNice []models.Article
	if models.Article1.NiceNum > 0 {
		articleWithNice = append(articleWithNice, models.Article1)
	}
	if models.Article2.NiceNum > 0 {
		articleWithNice = append(articleWithNice, models.Article2)
	}
	if(len(articleWithNice) > 0) {
		jsonData , err := json.MarshalIndent(articleWithNice  , "" , " ")
		if err != nil {
			http.Error(w , "failed to encode JSON\n" , http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	} else {
		io.WriteString(w , "No artilce with links found")
	}
}

// POST /comment のハンドラ
func PostComment(w http.ResponseWriter, trq *http.Request) {
	// io.WriteString(w, "Posting comments...\n")
	comments := models.AllComments
	jsonData , err := json.MarshalIndent(comments , "" , " ")
	if err != nil {
		http.Error(w , "faild to encode json\n" , http.StatusInternalServerError)
		return
	}
	w.Write(jsonData);
}

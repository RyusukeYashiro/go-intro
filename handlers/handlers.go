// Go では main 以外のパッケージ名は、基本的にそのファイルが格納されているディレクトリ名と
// 同名にする必
package handlers

import (
	"encoding/json"
	"go-intro/models"
	"go-intro/services"
	"io"
	"log"
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
	// var reqBodybuffer []byteの Read メソッドを呼び出すことで、リクエストボディの中身を引数に渡した
	// reqBodybuffer に読み出している
	// lenByte , err := strconv.Atoi(req.Header.Get("Content-Length"))
	// if err != nil {
	// 	http.Error(w , "cannot get content length\n" , http.StatusBadRequest)
	// 	return
	// }
	// reqBodybuffer := make([]byte , lenByte)
	// Readメゾットからは読み取り終わった時にio.EOFが帰ってくるこEOFが帰ってくることになる
	// if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err , io.EOF) {
		//Read メソッドからの err が io.EOF 以外だった場合、500 番エラーを返却
		// http.Error(w , "faield to get reqeust body\n" , http.StatusBadRequest)
		// return
	// }
	// defer req.Body.Close()

	var reqArticle models.Article
	// if err := json.Unmarshal(reqBodybuffer , &reqArticle); err != nil {
	// 	http.Error(w , "fail to decode json\n" , http.StatusBadRequest)
	// 	return
	// }

	// ストリームから直接データを受け取る
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w , "failed to decode json\n" , http.StatusBadRequest)
	}

	// データ取得→ json エンコード」	
	article , err := services.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w , "fail internal exec\n" , http.StatusInternalServerError)
		return
	}
	//　デーコードした内容を再度エンコードしてレスポンスする場合
	// article := reqArticle
	// jsonData , err := json.MarshalIndent(article , "" , " ")
	// if err != nil {
	// 	http.Error(w , "fail to encode json\n" , http.StatusInternalServerError)
	// 	return
	// }
	// http.ResponseWriterには []byte を引数にとって書き込み処理を行うことができる Write
	// メソッドが備わっている
	// w.Write(jsonData)	
	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Handling article list request")
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
	//jsonへのencode処理
	articleList , err := services.GetArticleListService(pg)
	if err != nil {
		log.Printf("Handler error: %v", err)  // エラーログを追加
		http.Error(w , "fail internal exec\n" , http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(articleList)	
}

// GET /article/{id} のハンドラ
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
	article , err := services.GetArticleService(artcile_Id)
	if err != nil {
		http.Error(w , "fail internal exec\n" , http.StatusInternalServerError)
		return
	}
	// jsonData , err := json.MarshalIndent(article , "" , " ")
	// if err != nil {
	// 	errMs := fmt.Sprintf("faild to encode json (articleID :%d)\n" , artcile_Id)
	// 	http.Error(w ,errMs  , http.StatusInternalServerError)
	// 	return
	// }
	json.NewEncoder(w).Encode(article)
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w , "Posting nice...\n")
	var articleWithNice models.Article
		// jsonData , err := json.MarshalIndent(articleWithNice  , "" , " ")
		// if err != nil {
		// 	http.Error(w , "failed to encode JSON\n" , http.StatusInternalServerError)
		// 	return
		// }
		// w.Write(jsonData)
	if err := json.NewDecoder(req.Body).Decode(&articleWithNice); err != nil {
		http.Error(w , "fail to decode json\n" , http.StatusBadRequest)
		return
	}
	article , err := services.PostNiceService(articleWithNice)
	if err != nil {
		http.Error(w  , "fail internal exec\n" , http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

// POST /comment のハンドラ
func PostComment(w http.ResponseWriter, req *http.Request) {
    // 正しい型を使用する
    var newComment models.Comment

    // リクエストボディをデコード
    if err := json.NewDecoder(req.Body).Decode(&newComment); err != nil {
        http.Error(w, "fail to decode json\n", http.StatusBadRequest)
        return
    }
	comment , err := services.PostCommentService(newComment)
	if err != nil {
		http.Error(w , "fail internal exec\n" , http.StatusInternalServerError)
		return
	}

    // 全コメントリストをレスポンスとして返す
    json.NewEncoder(w).Encode(comment)
}

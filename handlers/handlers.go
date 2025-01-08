// Go では main 以外のパッケージ名は、基本的にそのファイルが格納されているディレクトリ名と
// 同名にする必
package handlers

import (
	"encoding/json"
	"go-intro/models"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

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
	article := models.Article1
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

	log.Println(pg)
	//jsonへのencode処理
	articleList := []models.Article{models.Article1 , models.Article2}
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
	article := models.Article1
	// jsonData , err := json.MarshalIndent(article , "" , " ")
	// if err != nil {
	// 	errMs := fmt.Sprintf("faild to encode json (articleID :%d)\n" , artcile_Id)
	// 	http.Error(w ,errMs  , http.StatusInternalServerError)
	// 	return
	// }
	log.Println(artcile_Id)
	json.NewEncoder(w).Encode(article)
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
		// jsonData , err := json.MarshalIndent(articleWithNice  , "" , " ")
		// if err != nil {
		// 	http.Error(w , "failed to encode JSON\n" , http.StatusInternalServerError)
		// 	return
		// }
		// w.Write(jsonData)
		if err := json.NewDecoder(req.Body).Decode(&articleWithNice); err != nil {
			http.Error(w , "fail to decode json\n" , http.StatusBadRequest)
		}
		article := articleWithNice
		json.NewEncoder(w).Encode(article)
	} else {
		io.WriteString(w , "No artilce with links found")
	}
}

// POST /comment のハンドラ
func PostComment(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Posting comments...\n")

    // 正しい型を使用する
    var newComment models.Comment

    // リクエストボディをデコード
    if err := json.NewDecoder(req.Body).Decode(&newComment); err != nil {
        http.Error(w, "fail to decode json\n", http.StatusBadRequest)
        return
    }

    // コメントにIDや作成日時を付加
    newComment.CommentID = len(models.AllComments) + 1
    newComment.CreatedAt = time.Now()

    // 全コメントリストに新しいコメントを追加
    models.AllComments = append(models.AllComments, newComment)

    // 全コメントリストをレスポンスとして返す
    json.NewEncoder(w).Encode(models.AllComments)
}

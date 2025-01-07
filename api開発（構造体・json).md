構造体・jsonメモ

- POST /article 記事を投稿する 投稿に成功した記事の内容
- GET /article/list 記事一覧を取得する 記事のリスト
- GET /article/{id} 指定 ID の記事を取得する 記事の内容
- POST /article/nice 記事にいいねをつける いいねをつけた記事の内容
- POST /comment コメントを投稿する 投稿に成功したコメントの

*Comment 構造体の定義*
---------------------------------------------
• CommentID: コメント ID
• ArticleID: コメント対象となった記事 ID
• Message: コメント本文
• CreatedAt: 投稿日時

*記事構造体の定義*
---------------------------------------------
• 記事 ID: ブログサービスに投稿される全記事に振られる連番 -> ID フィールド
• 記事タイトル -> Title フィールド
• 記事本文 -> Contents フィールド
• 投稿者名 -> UserName フィールド
• いいね数 -> NiceNum フィールド
• コメント: その記事についたコメントをスライス形式で格納する-> CommentList フィールド
• 投稿日時 -> CreatedAt フィールド

 エンドポイントの設計
------------------------------------------------

エンドポイント          レスポンス内容 (json) 
POST /article           Article1 
GET /article/list       Article1 と Article2 のスライス
GET /article/{id}       Article1
POST /article/nice      Article1
POST /comment           Comment1


*デコード*
json エンコード同様、json から Go 構造体にデコードするための関数というのも encoding/json
パッケージの中に定義されています。それが json.Unmarshal 関数
```
func Unmarshal(data []byte, v any) error
```

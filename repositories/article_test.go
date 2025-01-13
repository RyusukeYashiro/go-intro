package repositories_test

// テストコードでは xxx_test というパッケージ名を使う方がいい

import (
	"database/sql"
	"fmt"
	"go-intro/models"
	"go-intro/repositories"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)	

	db , err := sql.Open("mysql" , dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	// 1. テスト結果にて期待する値を定義
	expected := models.Article{
		ID: 1,
		Title: "firstPost",
		Contents: "this is my first blog",
		UserName: "yashiro",
		NiceNum: 2,
	}
	

	// 2. テスト対象となる関数を実行
	get  , err := repositories.SelectArticleDetail(db , expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if get.ID != expected.ID {
		t.Errorf("ID: get %d but want %d\n" , get.ID , expected.ID)
	}
	if get.Title != expected.Title {
		t.Errorf("Title: get %s but want %s\n" , get.Title , expected.Title)
	}
	if get.Contents != expected.Contents {
		t.Errorf("Content: get %s but want %s\n", get.Contents, expected.Contents)
	}
	if get.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s\n", get.UserName, expected.UserName)
	}
	if get.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", get.NiceNum, expected.NiceNum)
	}
}
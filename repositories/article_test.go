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

	//ドリンブルンテストの実装
	tests := []struct {
		testTitle string
		expected models.Article
	}{
		{
			testTitle: "test1",
			expected : models.Article {
				ID: 1,
				Title: "firstPost",
				Contents: "this is my first blog",
				UserName: "yashiro",
				NiceNum: 2,
			},
		} , {
			testTitle: "test2",
			expected: models.Article {
				ID: 2,
				Title: "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum: 4,
			},
		},
	}

	for _, test := range tests {
		//個別のテストを回すための、単体テストを行うのには、runを用いる
		t.Run(test.testTitle , func(t *testing.T) {
			got , err := repositories.SelectArticleDetail(db , test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n" , got.ID , test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n" , got.Title , test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}
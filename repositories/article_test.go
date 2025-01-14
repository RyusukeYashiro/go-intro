package repositories_test

// テストコードでは xxx_test というパッケージ名を使う方がいい

import (
	"fmt"
	"go-intro/models"
	"go-intro/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleList(t *testing.T) {
	fmt.Println("TestSelectArticleList test....")
	expectedNum := 2
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestInsertArticle(t *testing.T) {
	fmt.Println("TestInsertArticle test....")
	article := models.Article {
		Title: "insertTest",
		Contents: "testest",
		UserName: "saki",
	}

	expectedArticleNum := 3
	newArticle , err := repositories.InsertArticle(testDB , article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
	}
	t.Cleanup(func() {
		const sqlStr = `delete from articles where title = ? and contents = ? and username = ?`
		testDB.Exec(sqlStr , article.Title , article.Contents , article.UserName)
	})
}

func TestAddNiceNum(t *testing.T) {
	fmt.Println("TestAddNiceNum test....");
	articleID := 1
	before , err := repositories.SelectArticleDetail(testDB , articleID)
	if err != nil {
		t.Fatal(err)
	}
	err = repositories.UpdateNiceNum(testDB , articleID)
	if err != nil {
		t.Fatal(err)
	}
	after , err := repositories.SelectArticleDetail(testDB , articleID)
	if err != nil {
		t.Fatal(err)
	}
	if after.NiceNum - before.NiceNum != 1 {
		t.Error("failed to update nice num")
	}
}

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	fmt.Println("TestSelectArticleDetail test....")
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
			got , err := repositories.SelectArticleDetail(testDB , test.expected.ID)
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
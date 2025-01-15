package repositories_test

import (
	"fmt"
	"go-intro/models"
	"go-intro/repositories"
	"testing"
)


func TestSelectCommentList(t *testing.T) {
	fmt.Println("TestSelectCommentList test....")
	articleID := 1
	got , err := repositories.SelectCommentList(testDB , articleID)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got ID %d\n", articleID, comment.ArticleID)
		}
	}
}

func TestInsertComment(t *testing.T) {
	fmt.Println("TestInsertComment test...")
	comment := models.Comment {
		ArticleID: 1,
		Message: "CommentInsertTest",
	}

	expectedCommentID := 3
	newComment , err := repositories.InsertComment(testDB , comment)
	if err != nil {
		t.Fatal(err)
		return
	}
	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n" , expectedCommentID , newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `delete from comments where message = ?`
		testDB.Exec(sqlStr , comment.Message)
	})

}
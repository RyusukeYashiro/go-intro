package services

import (
	"go-intro/models"
	"go-intro/repositories"
)

//コメントデータをデータベース内に挿入し、その値を返す
func PostCommentService (Comment models.Comment) (models.Comment , error) {
	db , err := connectDB()
	if err != nil {
		return models.Comment{} , err
	}
	defer db.Close()
	comment_posted , err := repositories.InsertComment(db , Comment)
	if err != nil {
		return models.Comment{} , err
	}
	return comment_posted , nil
}
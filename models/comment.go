package models

import (
	"JeffMusic/dao"
)

type CommentVO struct {
	Id       int    `json:"id"`
	Comment  string `json:"comment"`
	SongId   int    `json:"song_id"`
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Comment string `json:"comment"`
	SongId  int    `json:"song_id"`
	UserId  int    `json:"user_id"`
}

func CreateComment(comment *Comment) (err error) {
	err = dao.DB.Create(&comment).Error
	return
}

func GetComment(songId int) ([]CommentVO, error) {
	var comments []CommentVO
	if err := dao.DB.Debug().Select("id,comment,song_id,user_id,name as user_name").Table("comments").Joins("left outer join users "+
		"on comments.user_id = users.union_id where song_id=?", songId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

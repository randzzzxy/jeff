package models

import (
	"JeffMusic/dao"
)

type Reply struct {
	Id        int    `json:"id"`
	Reply     string `json:"reply"`
	CommentId int    `json:"comment_id"`
	UserId    int    `json:"user_id"`
}

func CreateReply(reply *Reply) (err error) {
	err = dao.DB.Create(&reply).Error
	return
}

func GetReply(commentId int) ([]Reply, error) {
	var replies []Reply
	if err := dao.DB.Debug().Where("comment_id=?", commentId).Find(&replies).Error; err != nil {
		return nil, err
	}
	return replies, nil
}

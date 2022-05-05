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

type ReplyVO struct {
	Id        int    `json:"id"`
	Reply     string `json:"reply"`
	CommentId int    `json:"comment_id"`
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
}

func CreateReply(reply *Reply) (err error) {
	err = dao.DB.Create(&reply).Error
	return
}

func GetReply(commentId int) ([]ReplyVO, error) {
	var replies []ReplyVO
	if err := dao.DB.Debug().Select("id,reply,comment_id,user_id,name as user_name").Table("replies").Joins("left outer join users "+
		"on replies.user_id = users.union_id where comment_id=?", commentId).Find(&replies).Error; err != nil {
		return nil, err
	}
	return replies, nil
}

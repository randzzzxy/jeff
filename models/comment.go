package models

import (
	"JeffMusic/dao"
	"time"
)

type Comment struct {
	Id        int       `json:"id"`
	Comment   string    `json:"comment"`
	SongId    int       `json:"song_id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateComment(comment *Comment) (err error) {
	err = dao.DB.Create(&comment).Error
	return
}

func GetComment(songId int) ([]Comment, error) {
	var comments []Comment
	if err := dao.DB.Debug().Where("song_id=?", songId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

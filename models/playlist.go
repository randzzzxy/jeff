package models

import (
	"JeffMusic/dao"
	"time"
)

type PlayList struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func CreatePlayList(playList *PlayList) (err error) {
	err = dao.DB.Create(&playList).Error
	return
}

func GetPlayList(userId int) ([]PlayList, error) {
	var playList []PlayList
	if err := dao.DB.Debug().Where("user_id=?", userId).Find(&playList).Error; err != nil {
		return nil, err
	}
	return playList, nil
}

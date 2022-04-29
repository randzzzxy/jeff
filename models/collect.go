package models

import (
	"JeffMusic/dao"
)

type Collect struct {
	Id         int `json:"id"`
	PlayListId int `json:"play_list_id"`
	SongId     int `json:"song_id"`
}

func CreateCollect(collect *Collect) (err error) {
	err = dao.DB.Create(&collect).Error
	return
}

func GetSongs(playListId int) ([]Song, error) {
	var songs []Song
	if err := dao.DB.Table("songs").Joins("inner join collects "+
		"on collects.song_id = songs.id where play_list_id=?", playListId).Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

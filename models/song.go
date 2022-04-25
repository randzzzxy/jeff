package models

import "JeffMusic/dao"

// Song Model
type Song struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	SongUrl   string `json:"song_url"`
	CoverUrl  string `json:"cover_url"`
	AuthorId  int    `json:"author_id"`
	LyricsUrl string `json:"lyrics_url"`
}

func UploadSong(song *Song) (err error) {
	err = dao.DB.Create(&song).Error
	return
}

package models

import (
	"JeffMusic/dao"
	"time"
)

// User Model
type User struct {
	UnionID        int       `json:"union_id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	AvatarUrl string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateNewAccount(user *User) (err error) {
	err = dao.DB.Create(&user).Error
	return
}

func GetUserInfo(id int) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Debug().Where("union_id=?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return
}

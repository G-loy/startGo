package models

import "github.com/jinzhu/gorm"

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func CheckUser(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, err
}
func GetRole(username, password string) (int, error) {
	var user User
	err := db.Select("role").Where(User{Username: username, Password: password}).First(&user).Error

	if err != nil {
		return user.Role, nil
	} else {
		return -1, err
	}

}

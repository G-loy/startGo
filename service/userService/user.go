package userService

import "firstGin/models"

type User struct {
	Username string
	Password string
}

func (a *User) Check() (bool, error) {
	return models.CheckUser(a.Username, a.Password)
}
func (a *User) GetRole() (int, error) {
	return models.GetRole(a.Username, a.Password)
}

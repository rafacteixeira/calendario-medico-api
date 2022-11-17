package database

import "github.com/rafacteixeira/calendario-medico-api/model"

func FindUser(login string) model.User {
	var obj model.User
	DB.Where("login = ?", login).First(&obj)
	return obj
}

func CreateUser(u *model.User) {
	DB.Create(&u)
}

func FindAdminRole() model.Role {
	adminRole := model.Role{
		Name: "admin",
	}
	DB.FirstOrCreate(&adminRole)
	return adminRole
}

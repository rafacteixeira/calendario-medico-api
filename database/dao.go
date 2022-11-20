package database

import "github.com/rafacteixeira/calendario-medico-api/model"

func FindUser(login string) model.User {
	var obj model.User
	DB.Where("login = ?", login).First(&obj)
	return obj
}

func FindUserWithRoles(login string) model.User {
	var obj model.User
	DB.Preload("Roles").Where("login = ?", login).First(&obj)
	return obj
}

func FindUserWithEventsAndNotes(id uint) model.User {
	var obj model.User
	DB.Preload("Events").Preload("Notes").Where("id = ?", id).First(&obj)
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

func CreateEvent(event model.Event) {
	DB.Create(&event)
}

func CreateNote(note model.Note) {
	DB.Create(&note)
}

func CreateToken(t model.Token) {
	DB.Create(&t)
}

func FindToken(jti string) model.Token {
	token := model.Token{}
	DB.Where("jti = ?", jti).Find(&token)
	return token
}

func DeleteToken(jti string) {
	DB.Where("jti = ?", jti).Delete(&model.Token{})
}

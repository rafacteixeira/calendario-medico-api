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

func CreateUser(u *model.User) {
	DB.Create(&u)
}

func CreateEvent(event model.Event) model.Event {
	dbObj := model.Event{}
	DB.Where(event).FirstOrCreate(&dbObj)
	return dbObj
}

func CreateNote(note model.Note) model.Note {
	dbObj := model.Note{}
	DB.Where(note).FirstOrCreate(&dbObj)
	return dbObj
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

func DeleteEvent(e model.Event) {
	DB.Where(e).Delete(&model.Event{})
}

func DeleteNote(n model.Note) {
	DB.Where(n).Delete(&model.Note{})
}

func FindEventsByUser(userId uint) []model.Event {
	var events []model.Event
	DB.Where("user_id = ?", userId).Find(&events)
	return events
}

func FindNotesByUser(userId uint) []model.Note {
	var notes []model.Note
	DB.Where("user_id = ?", userId).Find(&notes)
	return notes
}

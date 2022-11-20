package services

import (
	"github.com/rafacteixeira/calendario-medico-api/database"
	"github.com/rafacteixeira/calendario-medico-api/model"
)

var CreateEvent = database.CreateEvent
var CreateNote = database.CreateNote
var DeleteEvent = database.DeleteEvent
var DeleteNote = database.DeleteNote
var FindEventsByUser = database.FindEventsByUser
var FindNotesByUser = database.FindNotesByUser

func saveEvent(event model.Event, userID uint) model.Event {
	dbEvent := model.Event{
		Date:       event.Date,
		EventType:  event.EventType,
		EventWatch: event.EventWatch,
		UserID:     userID,
	}
	dbEvent = CreateEvent(dbEvent)
	return dbEvent
}

func saveNote(note model.Note, userID uint) model.Note {
	dbNote := model.Note{
		Date:   note.Date,
		Note:   note.Note,
		UserID: userID,
	}
	dbNote = CreateNote(dbNote)
	return dbNote
}

func SaveEvent(request model.Event, userID uint) model.Event {
	saved := saveEvent(request, userID)
	return saved
}

func SaveNote(request model.Note, userID uint) model.Note {
	saved := saveNote(request, userID)
	return saved
}

func RemoveEvent(event model.Event, userID uint) {
	event.UserID = userID
	DeleteEvent(event)
}

func RemoveNote(note model.Note, userID uint) {
	note.UserID = userID
	DeleteNote(note)
}

func ListEvents(userID uint) []model.Event {
	return FindEventsByUser(userID)
}

func ListNotes(userID uint) []model.Note {
	return FindNotesByUser(userID)
}

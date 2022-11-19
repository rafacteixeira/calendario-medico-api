package services

import (
	"github.com/rafacteixeira/calendario-medico-api/database"
	"github.com/rafacteixeira/calendario-medico-api/model"
	"github.com/rafacteixeira/calendario-medico-api/util"
)

var CreateEvent = database.CreateEvent
var CreateNote = database.CreateNote
var FindUserWithEventsAndNotes = database.FindUserWithEventsAndNotes

func SaveAllEvents(request util.EventsAndNotes, userID uint) error {

	events := request.Events
	saveEvents(events, userID)

	return nil
}

func SaveAllNotes(request util.EventsAndNotes, userID uint) error {

	notes := request.Notes
	saveNotes(notes, userID)

	return nil
}

func GetEventsAndNotesByUser(userID uint) util.EventsAndNotes {

	user := FindUserWithEventsAndNotes(userID)

	response := util.EventsAndNotes{
		Events: []model.Event{},
		Notes:  []model.Note{},
	}

	for _, event := range user.Events {
		response.Events = append(response.Events, event)
	}

	for _, note := range user.Notes {
		response.Notes = append(response.Notes, note)
	}

	return response
}

func saveEvents(events []model.Event, userID uint) {
	for _, event := range events {
		saveEvent(event, userID)
	}
}

func saveNotes(notes []model.Note, userID uint) {
	for _, note := range notes {
		saveNote(note, userID)
	}
}

func saveEvent(event model.Event, userID uint) {
	dbEvent := model.Event{
		Date:       event.Date,
		EventType:  event.EventType,
		EventWatch: event.EventWatch,
		UserID:     userID,
	}
	CreateEvent(dbEvent)
}

func saveNote(note model.Note, userID uint) model.Note {
	dbNote := model.Note{
		Date:   note.Date,
		Note:   note.Note,
		UserID: userID,
	}
	CreateNote(dbNote)
	return dbNote
}

func SaveEvent(request model.Event, userID uint) {
	saveEvent(request, userID)
}

func SaveNote(request model.Note, userID uint) model.Note {
	saved := saveNote(request, userID)
	return saved
}

func DeleteEvent(eventID uint, userID uint) {
}

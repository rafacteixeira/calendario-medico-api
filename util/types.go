package util

import (
	"github.com/rafacteixeira/calendario-medico-api/model"
)

type AuthRequest struct {
	Login    string
	Password string
}

type AuthError struct {
	Error   string
	Message string
}

type AuthResponse struct {
	Token string
}

type EventsAndNotes struct {
	Events []model.Event
	Notes  []model.Note
}

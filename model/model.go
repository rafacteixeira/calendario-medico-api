package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint    `gorm:"primaryKey"`
	Login    string  `gorm:"index"`
	Password string  `json:"-"`
	Roles    []Role  `json:",omitempty" gorm:"many2many:user_role;"`
	Events   []Event `json:",omitempty"`
	Notes    []Note  `json:",omitempty"`
}

type Role struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Users []User `gorm:"many2many:user_role;" json:",omitempty"`
}

type Event struct {
	gorm.Model
	Date       time.Time `json:"Date"`
	EventType  string    `json:"Type"`
	EventWatch string    `json:"Watch"`
	UserID     uint      `json:"-"`
}

type Note struct {
	gorm.Model
	Date   time.Time `json:"date"`
	Note   string    `json:"txt"`
	UserID uint      `json:"-"`
}
type ApiModel interface {
	User | Role | Note | Event
}

type Token struct {
	JTI       string         `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Token     string
}

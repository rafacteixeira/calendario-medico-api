package model

import "time"

type User struct {
	ID            uint   `gorm:"primaryKey"`
	Login         string `gorm:"index"`
	Password      string
	Roles         []Role         `json:",omitempty" gorm:"many2many:user_role;"`
	CalendarItems []CalendarItem `json:",omitempty"`
	Notes         []Note         `json:",omitempty"`
}

type Role struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Users []User `gorm:"many2many:user_role;" json:",omitempty"`
}

type ItemType string
type ItemWatch string

const (
	Enfermaria  ItemType = "Enf"
	Ambulatorio          = "Amb"
	Plantao              = "Pla"
	PosPlantao           = "PosP"
	Aula                 = "Aula"

	Manha ItemWatch = "manha"
	Tarde ItemWatch = "tarde"
	Noite ItemWatch = "noite"
)

type CalendarItem struct {
	ID        uint      `json:"ID"`
	Date      time.Time `json:"Date"`
	ItemType  ItemType  `json:"Type"`
	ItemWatch ItemWatch `json:"Watch"`
	UserID    uint      `json:"-"`
}

type Note struct {
	ID     uint      `json:"id"`
	Date   time.Time `json:"date"`
	Note   string    `json:"txt"`
	UserID uint      `json:"-"`
}

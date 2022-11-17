package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Login    string `gorm:"index"`
	Password string
	Roles    []Role `gorm:"many2many:user_role;"`
}

type Role struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Users []User `gorm:"many2many:user_role;"`
}

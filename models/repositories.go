package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Main struct {
	Book     Book     `gorm:"foreignKey:BookRefer"`
	Employee Employee `gorm:"foreignKey:EmployeeRefer"`
	Visitor  Visitor  `gorm:"foreignKey:VisitorRefer"`
}
type Book struct {
	Id    int64  `gorm:"primarykey" json:"id"`
	Name  string `gorm:"varchar(50)" json:"name"`
	Stock int32  `gorm:"integer(100)" json:"stock"`
}
type Employee struct {
	Id       int64  `gorm:"primarykey" json:"id"`
	Name     string `gorm:"varchar(50)" json:"name"`
	Position string `gorm:"varchar(50)" json:"position"`
	Sex      string `gorm:"varchar(1)" json:"sex"`
}
type Visitor struct {
	Id        int64  `gorm:"primarykey" json:"id"`
	Name      string `gorm:"varchar(50)" json:"name"`
	Status    string `gorm:"varchar(20)" json:"status"`
	Return    string `gorm:"type:date;default:null" json:"return"`
	CreatedAt time.Time
}
type User struct {
	Id       int64  `gorm:"primarykey" json:"id"`
	Fullname string `gorm:"varchar(50);column:full_name" json:"full_name"`
	Username string `gorm:"varchar(50)" json:"username"`
	Password string `gorm:"varchar(50)" json:"Password"`
	Email    string `gorm:"varchar(50)" json:"Email"`
}

func (user *User) SetPassword(password string) {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(hashedPass)
}

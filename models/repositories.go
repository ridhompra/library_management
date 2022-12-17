package models

import "time"

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
	Id         int64     `gorm:"primarykey" json:"id"`
	Name       string    `gorm:"varchar(50)" json:"name"`
	LoanDate   time.Time `gorm:"type:date" json:"loan_date"`
	Status     string    `gorm:"varchar(20)" json:"status"`
	ReturnDate time.Time `gorm:"type:date"  json:"return_date"`
}
type User struct {
	Id       int64  `gorm:"primarykey" json:"id"`
	Username string `gorm:"varchar(50)" json:"username"`
	Password string `gorm:"varchar(50)" json:"Password"`
	Email    string `gorm:"varchar(50)" json:"Email"`
}

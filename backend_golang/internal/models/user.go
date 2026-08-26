package models

import "time"

type User struct {
	ID				uint64		`gorm:"primaryKey"`
	RoleId			uint64		`gorm:"not null;index"`
	FullName		string 		`gorm:"size:50;not null"`
	Email			string		`gorm:"size:255;not null;uniqueIndex"`
	Passwordhash	string		`gorm:"size:255;not null"`
	CreatedAt 		time.Time	`gorm:"not null"`
	UpdatedAt		time.Time	`gorm:"not null"`

	Role Role `gorm:"foreignKey:RoleId"`
}
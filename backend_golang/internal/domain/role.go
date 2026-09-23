package domain

import "time"

type Role struct {
	ID        uint64	`gorm:"primaryKey"`
	Name      string 	`gorm:"size:50;not null;uniqueIndex"`
	CreatedAt time.Time	`gorm:"not null"`
	UpdatedAt time.Time	`gorm:"not null"`
}


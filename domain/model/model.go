package model

import "time"

type User struct {
	ID        uint64     `gorm:"column:id"`
	Email     string     `gorm:"column:email"`
	IDCardNo  string     `gorm:"column:idCardNo"`
	FirstName string     `gorm:"column:firstName"`
	LastName  string     `gorm:"column:lastName"`
	Address   string     `gorm:"column:address"`
	CreatedAt time.Time  `gorm:"column:createdAt"`
	UpdatedAt time.Time  `gorm:"column:updatedAt"`
	DeletedAt *time.Time `gorm:"column:deletedAt"`
}

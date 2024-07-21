package entities

import "github.com/jinzhu/gorm"

type Data struct {
	gorm.Model
	UserID uint32
	Key    string `gorm:"not null"`
	Type   string `gorm:"not null"`
	Data   []byte `gorm:"not null"`
}

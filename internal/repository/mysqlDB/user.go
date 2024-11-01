package mysqlDB

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UID      string `gorm:"not null"`
	Name     string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
}

func (*User) TableName() string {
	return "user"
}

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

type apiUser struct {
	UID   string
	Name  string
	Email string
}

func (u *User) ToApiUser() (*apiUser, error) {
	au := &apiUser{
		UID:   u.UID,
		Name:  u.Name,
		Email: u.Email,
	}

	return au, nil
}

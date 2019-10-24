package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);unique_index;not null"`
	Password string `gorm:"type:varchar(64);not null"`
	Phone    string `gorm:"type:char(11);unique_index;not null"`
}

type JwtToken struct {
	Token string `json:token`
}

func Do() {
	if !WriteDb.HasTable(&User{}) {
		if err := WriteDb.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}

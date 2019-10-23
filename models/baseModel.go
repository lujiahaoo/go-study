package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ReadDb *gorm.DB
var WriteDb *gorm.DB

func init() {
	var Rerr error
	var Werr error
	ReadDb, Rerr = gorm.Open("mysql", "root:123456@tcp(47.100.181.111:3341)/go?charset=utf8&parseTime=True&loc=Local")
	WriteDb, Werr = gorm.Open("mysql", "root:123456@tcp(47.100.181.111:3342)/go?charset=utf8&parseTime=True&loc=Local")
	if Rerr != nil {
		ReadDb.Close()
		WriteDb.Close()
		panic(Rerr)
	}
	if Werr != nil {
		ReadDb.Close()
		WriteDb.Close()
		panic(Werr)
	}
}

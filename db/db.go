package db

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	host     = "localhost"
	port     = "3306"
	username = "root"
	password = "toor"
	db       = "mitrais"
)

var (
	dbgorm *gorm.DB
	err    error
)

//ConfigureDB db configuration
func ConfigureDB() {
	dbgorm, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+db+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		log.Fatalln("Cannot connect to database")
	}

	dbgorm.SingularTable(true)
	dbgorm.DB().SetMaxIdleConns(10)
	dbgorm.DB().SetMaxOpenConns(30)
	dbgorm.DB().SetConnMaxLifetime(time.Minute)
}

//GetDB return dbgorm
func GetDB() *gorm.DB {
	return dbgorm
}

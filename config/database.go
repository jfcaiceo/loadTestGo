package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qor/validations"
)

type Database struct {
	Adapter  string
	Host     string
	Port     string
	Username string
	Password string
	Database string
	LogMode  bool
}

var db *gorm.DB

func DB(configDB map[string]string) *gorm.DB {
	database := configDB["name"]
	username := configDB["username"]
	password := configDB["password"]
	host := configDB["host"]
	port := configDB["port"]
	charset := configDB["charset"]
	logMode := configDB["logMode"]

	if db == nil {
		fmt.Println("new mysql database connection")
		temp, err := gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset="+charset+"&parseTime=True&loc=Local")
		validations.RegisterCallbacks(temp)

		mode, _ := strconv.ParseBool(logMode)

		temp.LogMode(mode)

		if err != nil {
			log.Fatalln(err)
		}

		db = temp
	}

	return db
}

package models

import (
	"firstGin/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.TYPE, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.USER,
		setting.DatabaseSetting.PASSWORD,
		setting.DatabaseSetting.HOST,
		setting.DatabaseSetting.Dbname))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

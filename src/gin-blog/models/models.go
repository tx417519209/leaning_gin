package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"xing.learning.gin/src/gin-blog/pkg/logging"
	"xing.learning.gin/src/gin-blog/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID         int   `gorm:"primary_key" json:"id"`
	CreatedOn  int64 `json:"created_on"`
	ModifiedOn int64 `json:"modified_on"`
}

func CloseDB() {
	sql, _ := db.DB()
	defer sql.Close()
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	if field := db.Statement.Schema.LookUpField("created_on"); field != nil {
		field.Set(db.Statement.ReflectValue, time.Now().Unix())
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if field := db.Statement.Schema.LookUpField("modified_on"); field != nil {
		now := time.Now().Unix()
		field.Set(db.Statement.ReflectValue, &now)
	}
}

func Setup() {
	if setting.DatabaseSetting.Type != "mysql" {
		logging.Fatal("only support mysql")
	}
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.DatabaseSetting.User, setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host, setting.DatabaseSetting.Name)
	db, err = gorm.Open(mysql.Open(dns))
	if err != nil {
		logging.Fatal(err)
	}
	db.Callback().Create().Before("*").Register("update_timestamp_create", updateTimeStampForCreateCallback)
	db.Callback().Update().Before("*").Register("update_timestamp_update", updateTimeStampForUpdateCallback)
	sql, _ := db.DB()
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Auth{})
	db.AutoMigrate(&Tag{})
}

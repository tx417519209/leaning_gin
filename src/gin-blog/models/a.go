package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"xing.learning.gin/src/gin-blog/pkg/setting"
)

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

func init() {
	if setting.DBType != "mysql" {
		log.Fatalf("only support mysql")
	}
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.DBUser, setting.DBPwd, setting.DBHost, setting.DBName)
	db, err = gorm.Open(mysql.Open(dns))
	if err != nil {
		log.Fatal(err)
	}
	db.Callback().Create().Before("*").Register("update_timestamp_create", updateTimeStampForCreateCallback)
	db.Callback().Update().Before("*").Register("update_timestamp_update", updateTimeStampForUpdateCallback)
	sql, _ := db.DB()
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
}

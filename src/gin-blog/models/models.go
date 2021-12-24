package models

import (
	"gorm.io/gorm"
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

package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Model(&Auth{}).Where(&Auth{UserName: username, Password: password}).First(&auth)
	return auth.ID > 0
}

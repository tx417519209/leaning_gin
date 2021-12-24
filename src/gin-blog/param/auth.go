package param

type AuthParam struct {
	UserName string `form:"username" json:"username" binding:"required,max=50"`
	PassWord string `form:"password" json:"password" binding:"required,max=50"`
}

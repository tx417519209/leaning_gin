package param

type ADDArticleParam struct {
	TagID     int    `form:"tag_id" json:"tag_id" binding:"required,gt=0"`
	Title     string `form:"title" json:"title" binding:"required,max=100"`
	Desc      string `form:"desc" json:"desc" binding:"required,max=255"`
	Content   string `form:"content" json:"content" binding:"required,max=65535"`
	CreatedBy string `form:"created_by" json:"created_by"  binding:"required"`
	State     int    `form:"state" json:"state" binding:"required,oneof=0 1"`
}

type GetArticleParam struct {
	ID int `uri:"article_id" binding:"required,gt=0"`
}

type GetArticlesParam struct {
	State *int `form:"state" json:"state" binding:"omitempty,oneof=0 1"`
	TagID *int `form:"tag_id" json:"tag_id" binding:"omitempty"`
}

type EditArticleParam struct {
	TagID      int    `form:"tag_id" json:"tag_id" binding:"required,gt=0"`
	Title      string `form:"title" json:"title" binding:"required"`
	Desc       string `form:"desc" json:"desc" binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"`
	ModifiedBy string `form:"modified_by" json:"modified_by"  binding:"required,max=100"`
	State      int    `form:"state" json:"state" binding:"required,oneof=0 1"`
}

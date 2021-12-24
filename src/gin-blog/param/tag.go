package param

type AddTagParam struct {
	Name      string `form:"name" json:"name" binding:"required"`
	State     int    `form:"state" json:"state" binding:"required,oneof=0 1"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,max=100"`
}

type EditTagParam struct {
	Name       *string `form:"name" json:"name" binding:"omitempty,max=20"`
	ModifiedBy string  `form:"modified_by" json:"modified_by" binding:"required,max=100"`
	State      *int    `form:"state" json:"state" binding:"omitempty,oneof=0 1"`
}

type GetTagsParam struct {
	Name  *string `form:"name" json:"name" binding:"omitempty"`
	State *int    `form:"state" json:"state" binding:"omitempty,oneof=0 1"`
}

type DeleteTagParam struct {
	ID int `uri:"id" json:"id" binding:"required,gte=0"`
}

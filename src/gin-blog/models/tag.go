package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreateBy   string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
func init() {
	db.AutoMigrate(&Tag{})
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	return tag.ID > 0
}

func AddTag(name string, state int, createBy string) bool {
	db.Create(&Tag{
		Name:     name,
		State:    state,
		CreateBy: createBy,
	})
	return true

}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID > 0
}
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, datas map[string]interface{}) bool {
	var tag Tag
	tag.ID = id
	db.Model(&tag).Save(datas)
	return true
}

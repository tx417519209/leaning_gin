package models

type Article struct {
	Model
	TagID      int    `json:"tag_id"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
func GetArticle(id int) (article Article) {
	db.Model(&Article{}).Where("id = ?", id).Find(&article)
	return
}
func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func init() {
	db.AutoMigrate(&Article{})
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true

}

func ExistArticleByID(id int) bool {
	var article Article
	db.Model(&Article{}).Select("id").Where("id = ?", id).First(&article)
	return article.ID > 0
}
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}

func EditArticle(id int, datas map[string]interface{}) bool {
	var article Article
	article.ID = id
	db.Model(&article).Updates(datas)
	return true
}

func init() {
	db.AutoMigrate(&Article{})
}

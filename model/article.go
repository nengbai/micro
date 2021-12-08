package model

// type Article struct {
// 	ArticleId uint64 `gorm:"column:articleId" json:"articleId"` // 自增
// 	Subject   string `gorm:"column:subject" json:"title"`       //
// 	Url       string `gorm:"column:url" json:"url"`
// 	ImgUrl    string `json:"imgurl"`
// 	HeadUrl   string `json:"headurl"`
// }

type ArticleBase struct {
	ArticleId uint64 `gorm:"column:articleId" json:"articleId"` // 自增
	Subject   string `gorm:"column:subject" json:"title"`       //
	Url       string `gorm:"column:url" json:"url"`
}

type Article struct {
	ArticleBase
	ImgUrl  string `json:"imgurl"`
	HeadUrl string `json:"headurl"`
}

func (ArticleBase) TableName() string {
	return "article"
}

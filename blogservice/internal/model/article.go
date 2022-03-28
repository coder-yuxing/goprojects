package model

type Article struct {
	Id            uint32 `gorm:"primary_key" json:"id"` // id
	Title         string `json:"title"`                 // 文章标题
	Desc          string `json:"desc"`                  // 文章简述
	CoverImageUrl string `json:"cover_image_url"`       // 封面图片地址
	Content       string `json:"content"`               // 文章内容
	State         uint8  `json:"state"`                 // 状态 0 为禁用、1 为启用
	*Model
}

func (model Article) TableName() string {
	return "blog_article"
}

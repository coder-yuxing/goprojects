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

// IArticleRepository 文章仓储接口
type IArticleRepository interface {
	Create(*Article) error           // 新建文章
	Update(*Article) error           // 更新文章
	Delete(uint32) error             // 删除文章
	GetById(uint32) (Article, error) // 查询指定文章
	GetAll() ([]Article, error)      // 查询全部文章
}

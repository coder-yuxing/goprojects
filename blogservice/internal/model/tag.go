package model

type ArticleTag struct {
	Id    uint32 `gorm:"primary_key" json:"id"` // id
	Name  string `json:"name"`                  // 标签名称
	State uint8  `json:"state"`                 // 状态 0 为禁用、1 为启用
	*Model
}

func (model ArticleTag) TableName() string {
	return "blog_tag"
}

// ArticleTagRepository 文章标签仓储接口
type ArticleTagRepository interface {
	Create(*ArticleTag) error           // 新建标签
	Update(*ArticleTag) error           // 更新标签
	Delete(uint32) error                // 删除标签
	GetById(uint32) (ArticleTag, error) // 查询指定标签
	GetAll() ([]ArticleTag, error)      // 查询全部标签
}

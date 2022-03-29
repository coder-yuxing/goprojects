package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yuxing/goprojects/blogservice/global"
	"github.com/yuxing/goprojects/blogservice/internal/model"
)

type ArticleTagRepository struct {
	engine *gorm.DB
}

func (r *ArticleTagRepository) Create(a *model.ArticleTag) error {
	return r.engine.Create(&a).Error
}

func (r *ArticleTagRepository) Update(a *model.ArticleTag) error {
	return r.engine.Model(&a).Where("id = ?", a.Id).Update(a).Error
}

func (r *ArticleTagRepository) Delete(id uint32) error {
	return r.engine.Delete(&model.ArticleTag{}, id).Error
}

func (r *ArticleTagRepository) GetById(id uint32) (model.ArticleTag, error) {
	tag := model.ArticleTag{}
	err := r.engine.First(&tag, id).Error
	if err != nil {
		return tag, err
	}
	return tag, nil
}

func (r *ArticleTagRepository) GetAll() ([]model.ArticleTag, error) {
	var tags []model.ArticleTag
	err := r.engine.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func NewTagRepository() ArticleTagRepository {
	return ArticleTagRepository{
		engine: global.DBEngine,
	}
}

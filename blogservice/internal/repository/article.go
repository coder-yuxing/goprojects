package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yuxing/goprojects/blogservice/internal/model"
)

type ArticleRepository struct {
	engine *gorm.DB
}

func (r *ArticleRepository) Create(a *model.Article) error {
	return r.engine.Create(&a).Error
}

func (r *ArticleRepository) Update(a *model.Article) error {
	return r.engine.Model(&a).Where("id = ?", a.Id).Update(a).Error
}

func (r *ArticleRepository) Delete(id uint32) error {
	return r.engine.Delete(&model.Article{}, id).Error
}

func (r *ArticleRepository) GetById(id uint32) (model.Article, error) {
	tag := model.Article{}
	err := r.engine.First(&tag, id).Error
	if err != nil {
		return tag, err
	}
	return tag, nil
}

func (r *ArticleRepository) GetAll() ([]model.Article, error) {
	var tags []model.Article
	err := r.engine.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

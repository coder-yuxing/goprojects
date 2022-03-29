package service

import (
	"github.com/yuxing/goprojects/blogservice/internal/model"
	"github.com/yuxing/goprojects/blogservice/internal/repository"
)

type ArticleTagService struct {
	repos repository.ArticleTagRepository
}

type CountTagCommand struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListCommand struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagCommand struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagCommand struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagCommand struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *ArticleTagService) Create(command CreateTagCommand) error {
	tag := model.ArticleTag{
		Name:  command.Name,
		State: command.State,
	}
	tag.CreatedBy = command.CreatedBy
	return svc.repos.Create(&tag)
}

func (svc *ArticleTagService) Update(command UpdateTagCommand) error {
	tag := model.ArticleTag{Id: command.ID, Name: command.Name, State: command.State}
	tag.ModifiedBy = command.ModifiedBy
	return svc.repos.Update(&tag)
}

func (svc *ArticleTagService) Delete(command DeleteTagCommand) error {
	return svc.repos.Delete(command.ID)
}

func (svc *ArticleTagService) GetById(id uint32) (model.ArticleTag, error) {
	return svc.repos.GetById(id)
}

func (svc *ArticleTagService) GetAll() ([]model.ArticleTag, error) {
	return svc.repos.GetAll()
}

func NewTagService() *ArticleTagService {
	return &ArticleTagService{
		repos: repository.NewTagRepository(),
	}
}
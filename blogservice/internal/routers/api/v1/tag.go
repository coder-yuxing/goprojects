package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yuxing/goprojects/blogservice/global"
	"github.com/yuxing/goprojects/blogservice/internal/service"
	"github.com/yuxing/goprojects/blogservice/pkg/app"
)

type Tag struct {

}

func NewTag() Tag {
	return Tag{}
}

// Create
// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	command := service.CreateTagCommand{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &command)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
	}

	err := service.NewTagService().Create(command)
	if err != nil {
		global.Logger.Errorf(c, "tagService.Create err: %v", err)
	}
	response.ToResponse(gin.H{})
}

// Package routers 路由管理
package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.New()
	// 配置共用组件
	r.Use(gin.Logger())   // 日志
	r.Use(gin.Recovery()) // 异常处理

	// 配置分组路由
	apiV1 := r.Group("/api/v1")
	{
		// 标签管理相关接口
		apiV1.POST("/tags")            // 新增标签
		apiV1.DELETE("/tags/:id")      // 删除指定标签
		apiV1.PUT("/tags/:id")         // 更新指定标签
		apiV1.GET("/tags")             // 获取标签列表
		apiV1.PATCH("/tags/:id/state") // 更新指定标签状态

		// 文章管理相关接口
		apiV1.POST("/articles")            // 新增文章
		apiV1.DELETE("/articles/:id")      // 删除指定文章
		apiV1.PUT("/articles/:id")         // 更新指定文章
		apiV1.PATCH("/articles/:id/state") // 更新指定文章状态
		apiV1.GET("/articles/:id")         // 获取指定文章
		apiV1.GET("/articles")             // 获取文章列表
	}
	return r
}

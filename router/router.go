package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gin_forum/controllers"
	"gin_forum/router/middleware"
	"gin_forum/config/logger"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 添加中间件
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(mw...)

	// 404 处理
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "该路由不存在"})
	})

	// 注册路由
	g.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	v1 := g.Group("/api/v1")
	v1.POST("/register", controllers.Register)
	v1.POST("/login", controllers.Login)

	v1.Use(middleware.AuthMiddleware())  // 增加token校验中间件
	{
		v1.GET("/category", controllers.CategoryList)
		v1.GET("/category/:id", controllers.CategoryDetail)
	}

	return g
}

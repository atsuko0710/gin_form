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
	g.POST("/register", controllers.Register)
	g.POST("/login", controllers.Login)
	
	return g
}

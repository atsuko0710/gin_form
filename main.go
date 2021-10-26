package main

import (
	"fmt"
	"gin_forum/config"
	"gin_forum/config/mysql"
	"gin_forum/config/redis"
	"gin_forum/config/logger"
	"gin_forum/controllers"
	"gin_forum/pkg/snowflake"
	"gin_forum/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	// 加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	
	if err := logger.Init(); err != nil {
		fmt.Printf("load logger config failed, err:%v\n", err)
		return
	}

	// 初始化数据库
	if err := mysql.Init(); err != nil {
		fmt.Printf("load mysql config failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	if err := snowflake.Init("2021-10-11", 1); err != nil {
		fmt.Printf("init snowflake failed, err:%v", err)
		return
	}

	// 初始化验证器的翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init validator failed, err:%v", err)
		return
	}

	// 初始化redis
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err:%v", err)
		return
	}
	defer redis.Clone()

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	router.Load(
		g,
	)

	http.ListenAndServe(config.Conf.Addr, g).Error()
}

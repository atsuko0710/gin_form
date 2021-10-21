package main

import (
	"fmt"
	"gin_forum/config"
	"gin_forum/config/mysql"
	"gin_forum/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	// 加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
	}
	if err := mysql.Init(); err != nil {
		fmt.Printf("load mysql config failed, err:%v\n", err)
	}
	defer mysql.Close()

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	router.Load(
		g,
	)

	http.ListenAndServe(config.Conf.Addr, g).Error()
}

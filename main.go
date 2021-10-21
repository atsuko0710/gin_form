package main

import (
	"gin_forum/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")
	g := gin.New()

	router.Load(
		g,
	)

	http.ListenAndServe(":8080", g).Error()
}

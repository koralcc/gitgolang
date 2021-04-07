package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	//consul服务注册
	ginRouter := gin.Default()
	data := make([]interface{}, 0)
	ginRouter.Handle("GET", "/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data": data,
		})
	})

	server := web.NewService(
		web.Address(":8000"),
		web.Handler(ginRouter),
	)
	server.Run()
}

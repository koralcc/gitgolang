package main

import (
	help "hello/Helper"
	"hello/ProdService"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	//consul服务注册
	consulreg := consul.NewRegistry(
		registry.Addrs("192.168.238.131:8500"), //consul服务部署的地址
	)

	//gin路由
	ginRouter := gin.Default()
	vlGroup := ginRouter.Group("/v1")
	{
		vlGroup.Handle("POST", "/prods", func(context *gin.Context) {
			var pr help.ProdsRequest
			err := context.Bind(&pr)
			if err != nil || pr.Size <= 0 {
				pr = help.ProdsRequest{
					2,
				}
			}
			context.JSON(200,
				gin.H{
					"data": ProdService.NewProdList(pr.Size)})
		})
	}

	//将gin整合到go-micro
	server := web.NewService(
		web.Name("prodservice"),
		//web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(consulreg), //注册服务到consul平台
	)
	server.Init()
	server.Run()

}

//用docker启动一个服务【最简单的容器启动】
// 1.拉取镜像(这里使用单机)
// docker pull consul
// 2.启动一个服务端(为了演示我们只启动一个)
// sudo docker run -d --name=cs -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0
// -d:后台执行
// -server：代表以服务端的方式启动
// -bootstrap：指定自己为leader，而不需要选举
// -ui：启动一个内置管理web界面
// -client：指定客户端可以访问的IP，设置为0.0.0.0则任意访问，否则默认本机可以访问
// （8500是后台UI端口，别忘了sudo iptable -l INPUT -p tcp --dport 8500 -j ACCEPT）
// 3.接下来就可以访问http://ip:8500/查看后台情况

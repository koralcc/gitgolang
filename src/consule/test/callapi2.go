package main

import (
	"context"
	"fmt"
	Models "hello/Models/protos"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	//服务发现
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.238.131:8500"),
	)

	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.Random),
	)
	CallAPI2(mySelector)
}

//client/http调用
func CallAPI2(s selector.Selector) {
	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("prodservice", "/v1/prods",
		//map[string]int{"size": 4})
		Models.ProdsRequest{
			Size: 3,
		})
	var rsp Models.ProdListResponse

	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.GetData())
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

//安装protobuf'等三个工具
// https://github.com/protocolbuffers/protobuf/releases //protoc.exe
// go get github.com/golang/protobuf/protoc-gen-go //protoc-gen-go.exe，配合protoc.exe生成go文件
// go get github.com/golang/protobuf/proto //proto,生成的go文件依赖包
// go get github.com/micro/protoc-gen-micro//go-micro 以来文件
// 看到如下三个说明成功
// protoc.exe
// protoc-gen-go.exe
// protoc-gen-micro.exe
//一个第三方库，用来处理protobuf生成的结构体，tag问题
//go get -u github.com/favadi/protoc-go-inject-tag
//   inject包导入后，加两部
//   1.proto文件加注释@inject_tag:json:"pName"
//   2.protoc执行加命令行protoc-go-inject-tag -input=../Prods.pb.go

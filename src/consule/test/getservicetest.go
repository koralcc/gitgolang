package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.238.131:8500"),
	)
	getService, err := consulReg.GetService("prodservice") //获取服务列表
	if err != nil {
		log.Fatal(err)
	}
	//next := selector.Random(getService) //随机选择一个服务
	for {
		next := selector.RoundRobin(getService) //轮询
		node, err := next()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(node.Address)
		time.Sleep(time.Second)
	}
}

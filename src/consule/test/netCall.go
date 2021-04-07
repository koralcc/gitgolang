package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	next := selector.RoundRobin(getService) //轮询
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("node:", node)
	callRes, err := CallAPI(node.Address, "/v1/prods", "POST")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("res:", callRes)
}

//原生调用
func CallAPI(addr string, path string, method string) (string, error) {
	req, err := http.NewRequest(method, "http://"+addr+path, nil)
	if err != nil {
		fmt.Println("构造request err：", err)
	}
	fmt.Println(req, "req")
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

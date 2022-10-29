package main

import (
	"github.com/julienschmidt/httprouter"
	"go-micro/web"
	"log"
)

func main() {
	//搭建web服务
	service := web.NewService(
		web.Name("go.micro.web.BikeWeb"),
		web.Version("latest"),
		web.Address(":10086"),
	)
	//服务初始化
	err := service.Init()
	if err != nil {
		log.Fatal(err)
	}
	//路由
	rou := httprouter.New()
	service.Handle("/", rou)
	//服务运行
	err1 := service.Run()
	if err1 != nil {
		log.Fatal(err1)
	}

}

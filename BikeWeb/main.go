package main

import (
	"BikeWeb/handler"
	"github.com/julienschmidt/httprouter"
	"go-micro/web"
	"log"
	"net/http"
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
	rou.NotFound = http.FileServer(http.Dir("html"))
	service.Handle("/", rou)
	rou.GET("/api/v1.0/areas", handler.GetArea)
	rou.GET("/api/v1.0/house/index", handler.GetIndex)
	rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCd)
	rou.GET("/api/v1.0/users", handler.PostRet)
	rou.GET("api/v1.0/sessions", handler.PostLogin)
	rou.GET("/api/v1.0/session", handler.DeleteSession)

	//服务运行
	err1 := service.Run()
	if err1 != nil {
		log.Fatal(err1)
	}

}

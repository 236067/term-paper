package handler

import (
	"BikeWeb/model"
	"BikeWeb/utils"
	"beego-develop/adapter/orm"
	"context"
	"fmt"
	"log"
	"time"

	example "PostRet/proto/example"
)

type Example struct{}

func (e *Example) PostRet(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println(" 注册服务  PostRet  /api/v1.0/users")
	//1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	/****2.连接redis**/
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis连接失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/****5.对接收到的密码进行加密**/
	user := model.User{}
	user.Password_hash = utils.Getmd5string(req.Password)
	user.Mobile = req.Mobile
	user.Name = req.Mobile
	/****6.插入数据到数据库中**/
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		fmt.Println("用户数据注册插入失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	/****7.生成sessionid**/
	sessionid := utils.Getmd5string(req.Mobile + req.Password)
	/****8.通过sessionid将数据返回redis**/
	rsp.Sessionid = sessionid
	bm.Put(sessionid+"user_id", id, time.Second*600)
	bm.Put(sessionid+"mobile", user.Mobile, time.Second*600)
	bm.Put(sessionid+"name", user.Name, time.Second*600)
	return nil
}

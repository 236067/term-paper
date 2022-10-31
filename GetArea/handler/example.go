package handler

import (
	"BikeWeb/model"
	"BikeWeb/utils"
	example "GetArea/proto/example"
	"context"
	"fmt"
	"github.com/astaxie/beego/client/orm"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取地域信息服务   GetArea  /api/v1.0/areas")
	//1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	//2.从缓存中获取数据

	//redis key

	//接收数据
	var areas []model.Area
	//获取redis中的数据

	//3.从数据库获取数据
	o := orm.NewOrm()
	//设置查询条件
	qs := o.QueryTable("area")
	//查询全部
	num, err := qs.All(&areas)

	if err != nil {
		fmt.Println("查询数据库失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
	}
	if num == 0 {
		fmt.Println("数据库为空")
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(utils.RECODE_NODATA)
	}
	//4.将查询到的数据存入redis中

	//5.将查询到的数据传入protobuf中
	for _, value := range areas {
		temp := example.Response_Address{}
		temp.Aid = int32(value.Id)
		temp.Aname = value.Name
		rsp.Data = append(rsp.Data, &temp)
	}
	return nil

}

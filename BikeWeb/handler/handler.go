package handler

import (
	DELETESESSION "DeleteSession/proto/example"
	GETAREA "GetArea/proto/example"
	GETIMAGECD "GetImageCd/proto/example"
	GETSESSION "GetSession/proto/example"
	POSTAVATAR "PostAvatar/proto/example"
	"go-micro.dev/v4"

	GETUSERINFO "GetUserInfo/proto/example"
	"BikeWeb/model"
	"BikeWeb/utils"

	POSTLOGIN "PostLogin/proto/example"
	POSTRET "PostRet/proto/example"

	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	beego "github.com/astaxie/beego/adapter"
	"github.com/julienschmidt/httprouter"
	//"github.com/micro/go-micro/client"

	"image"
	"image/png"
	"log"
	"net/http"
)

func GetArea(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode the incoming request as json	beego.Info("获取地区请求客户端 url：api/v1.0/areas")
	cli := micro.NewService()

	cli.Init()
	// call the backend service
	exampleClient := GETAREA.NewExampleService("", cli.Client())

	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//接收数据
	var areas []models.Area
	for _, value := range rsp.Data {
		temp := models.Area{Id: int(value.Aid), Name: value.Aname}
		areas = append(areas, temp)
	}
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   areas,
	}

	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func GetIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  "0",
		"errmsg": "ok",
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetImageCd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取图片验证码 url：/api/v1.0/imagecode/:uuid")
	uuid := ps.ByName("uuid")
	cli := micro.NewService()
	cli.Init()
	// call the backend service
	exampleClient := GETIMAGECD.NewExampleService("go.micro.srv.GetImageCd", cli.Client())
	rsp, err := exampleClient.GetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid: uuid,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//判断是否返回图片
	if rsp.Errno != "0" {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"errno":  rsp.Errno,
			"errmsg": rsp.Errmsg,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	//拼接图片结构体发送给前端
	var img image.RGBA
	for _, value := range rsp.Pix {
		img.Pix = append(img.Pix, uint8(value))
	}
	img.Stride = int(rsp.Stride)
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)

	var image captcha.Image
	image.RGBA = &img
	fmt.Println(image)
	png.Encode(w, image)
}

func PostRet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" 注册服务  PostRet  /api/v1.0/users")
	//接受 前端发送过来数据的
	var request map[string]interface{}
	// 将前端 json 数据解析到 map当中
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if request["mobile"].(string) == "" || request["password"].(string) == "" || request["sms_code"].(string) == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	//创建 grpc 客户端
	cli := micro.NewService()
	//客户端初始化
	cli.Init()

	//通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := POSTRET.NewExampleService("go.micro.srv.PostRet", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.PostRet(context.TODO(), &POSTRET.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
		SmsCode:  request["sms_code"].(string),
	})

	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//设置cookie
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{Name: "userlogin", Value: rsp.Sessionid, MaxAge: 600, Path: "/"}
		http.SetCookie(w, &cookie)
	}
	//将数据返回前端
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func PostLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("登陆 api/v1.0/sessions")
	//接受 前端发送过来数据的
	var request map[string]interface{}
	// 将前端 json 数据解析到 map当中
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//校验数据
	if request["mobile"].(string) == "" || request["password"].(string) == "" {
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// encode and write the response as json
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		return
	}
	//创建 grpc 客户端
	cli := micro.NewService()
	//客户端初始化
	cli.Init()

	//通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := POSTLOGIN.NewExampleService("go.micro.srv.PostLogin", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.PostLogin(context.TODO(), &POSTLOGIN.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
	})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	cookie, err := r.Cookie("userlogin")
	log.Println(cookie)
	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{Name: "userlogin", Value: rsp.Sessionid, MaxAge: 400, Path: "/"}
		http.SetCookie(w, &cookie)
	}
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func DeleteSession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode the incoming request as json
	beego.Info("退出登陆 url：/api/v1.0/session Deletesession()")

	cli := micro.NewService()
	cli.Init()
	// call the backend service
	exampleClient := DELETESESSION.NewExampleService("go.micro.srv.DeleteSession", cli.Client())
	//获取session
	userlogin, err := r.Cookie("userlogin")
	if err != nil || userlogin.Value == "" {
		log.Println("user not login")
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	rsp, err := exampleClient.DeleteSession(context.TODO(), &DELETESESSION.Request{
		Sessionid: userlogin.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if rsp.Errno == "0" {
		//将cookie中的sessionid设置为空
		_, err = r.Cookie("userlogin")
		if err == nil {
			cookie := http.Cookie{Name: "Iuserlogin", Path: "/", MaxAge: -1}
			http.SetCookie(w, &cookie)
		}
	}
	//返回数据
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	//设置格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("获取用户信息 GetUserInfo /api/v1.0/user")
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//创建 grpc 客户端
	cli := micro.NewService()
	//客户端初始化
	cli.Init()

	//通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := GETUSERINFO.NewExampleService("go.micro.srv.GetUserInfo", cli.Client())

	//通过句柄调用服务端函数
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &GETUSERINFO.Request{
		Sessionid: cookie.Value,
	})

	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//准备返回数据
	data := make(map[string]interface{})
	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	// encode and write the response as json
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//上传用户头像
func PostAvatar(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	beego.Info("上传用户touxiang PostAvatar /api/v1.0/user/avatar")
	//获取sessionid
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	file, header, err := r.FormFile("avatar")
	if err != nil {
		beego.Info("get file err:", err)
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	//文件校验
	filebuffer := make([]byte, header.Size)
	_, err = file.Read(filebuffer)
	if err != nil {
		beego.Info("get file err:", err)
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	//创建 micro 客户端
	cli := micro.NewService()
	//客户端初始化
	cli.Init()

	//通过protobuf生成文件创建连接服务端的客户端句柄
	exampleClient := POSTAVATAR.NewExampleService("go.micro.srv.PostAvatar", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.PostAvatar(context.TODO(), &POSTAVATAR.Request{
		Filename:  header.Filename,
		Filesize:  header.Size,
		Sessionid: cookie.Value,
		Avatar:    filebuffer, //file文件
	})

	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//返回数据
	data := make(map[string]interface{})
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	log.Println("data is ", data)
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetUserAuth(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("获取用户信息 GetUserInfo /api/v1.0/user")
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//创建 micro 客户端
	cli := micro.NewService()
	//客户端初始化
	cli.Init()

	//通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := GETUSERINFO.NewExampleService("go.micro.srv.GetUserInfo", cli.Client())

	//通过句柄调用服务端函数
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &GETUSERINFO.Request{
		Sessionid: cookie.Value,
	})

	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//准备返回数据
	data := make(map[string]interface{})
	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	// encode and write the response as json
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

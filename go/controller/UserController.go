package controller

import "C"
import (
	"encoding/json"
	"git.zx-tech.net/ljhua/hst"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"simpleMVC/go/entity"
	"simpleMVC/go/service"
)

//新增用户
func CreateController(c *hst.Context) {
	var user = entity.User{}

	r := c.R.Body
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("读取请求中数据失败")
	}

	err1 := json.Unmarshal(data, &user)
	if err1 != nil {
		log.Fatal("解析数据到结构体失败", err1)
		return
	}

	err2 := service.CreateUser(&user)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, struct {
			error error
		}{error: err2})
	} else {
		c.JSON(http.StatusOK, hst.JSONData{
			No: 0,
			Data: struct {
				msg  string
				data entity.User
			}{msg: "success",
				data: user},
		})
	}

	////日志输出
	//c.Str("llll", "awaa")

}

//查询所有用户
func GetUserList(c *hst.Context) {
	//计算耗时
	f := hst.AggHandleTime(c)
	defer f()

	result, err := service.GetAllUser()
	if err != nil {
		c.JSON2(http.StatusBadRequest, 0, err)
	} else {
		c.JSON2(http.StatusOK, 0, result)
	}
}

func Login(c *hst.Context) {
	if c.R.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(c.W.ResponseWriter, nil)
	} else {
		c.R.ParseForm()

		username := c.R.FormValue("username")
		password := c.R.FormValue("password")

		quesUser, err := service.GetUserByName(username)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		if quesUser.Password != password {
			c.JSON(http.StatusBadRequest, "用户密码错误")
		}

		//登录操作成功 用户信息写入session 返回给客户端
		c.SessionSet("userInfo", quesUser)

		c.JSON(http.StatusOK, struct { //常见错误
			Id   int
			Name string
			Msg  string
		}{
			Id:   quesUser.Id,
			Name: quesUser.Name,
			Msg:  "登录成功",
		})
	}
}

package controller

import "C"
import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"git.zx-tech.net/ljhua/hst"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"simpleMVC/go/entity"
	"simpleMVC/go/service"
)

var salt1 = "@#$%"
var salt2 = "^&*()"

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

	//判断邮箱是否合理
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, user.Email); !m {
		c.JSON(http.StatusBadRequest, "邮箱格式有误")
	}

	//判断该用户名是否已经注册
	user1, _ := service.GetUserByName(user.Name)
	if user1 != nil {
		c.JSON(http.StatusBadRequest, "该用户名已经被注册")
	}

	// 将密码加密存入数据库
	pwdMd5 := encryption(user.Password)
	user.Password = pwdMd5

	err2 := service.CreateUser(&user)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, struct {
			error error
		}{error: err2})
	} else {
		c.JSON(http.StatusOK, hst.JSONData{
			No: 0,
			Data: struct {
				Msg  string
				Data entity.User
			}{Msg: "success",
				Data: user},
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
		//c.R.ParseForm()
		//username := c.R.Form["username"]

		username := c.R.FormValue("username") //调用formvalue会默认提前调用parseform 但是只会返回同名参数的第一个
		password := c.R.FormValue("password")

		quesUser, err := service.GetUserByName(username)
		if quesUser == nil {
			c.JSON(http.StatusBadRequest, "未找到该用户")
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		pwdMD5 := encryption(password)
		if quesUser.Password != pwdMD5 {
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

//将密码进行MD5加密
func encryption(pwd string) string {
	h := md5.New()

	//将密码进行加密
	io.WriteString(h, pwd)
	pwdmd5 := fmt.Sprintf("%x", h.Sum(nil))

	//对密码加盐进行进一步加密
	io.WriteString(h, salt1)
	io.WriteString(h, pwdmd5)
	io.WriteString(h, salt2)

	lastPwd := fmt.Sprintf("%x", h.Sum(nil))
	return lastPwd
}

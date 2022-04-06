package controller

import (
	"git.zx-tech.net/ljhua/hst"
	"net/http"
	"simpleMVC/go/service"
	"strconv"
)

type User struct{}

func (user *User) GetAllUsers(c *hst.Context) {
	result, err := service.GetAllUser()
	if err != nil {
		c.JSON2(http.StatusBadRequest, 400, err)
	} else {
		c.JSON2(http.StatusOK, 200, result)
	}
}

func (user *User) LoginCheck(c *hst.Context) {

	c.R.ParseForm()
	username := c.R.FormValue("name")
	password := c.R.FormValue("password")

	quesUser, err := service.GetUserByName(username)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	if quesUser.Password != password {
		c.JSON(http.StatusBadRequest, "用户密码错误")
	}

	//登录操作成功 用户信息写入session 返回给客户端
	c.SessionSet("userInfo", user)

	c.JSON(http.StatusOK, struct { //常见错误
		Id   int
		Name string
		Msg  string
	}{
		Id:   quesUser.Id,
		Name: quesUser.Name,
		Msg:  "登录成功",
	})

	//c.JSON2(http.StatusOK, 0, quesUser)
}

func (user *User) GetUserById(c *hst.Context) {
	c.R.ParseForm()
	id := c.R.FormValue("id")
	id1, _ := strconv.Atoi(id)
	user1, err := service.GetUserById(id1)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, hst.JSONData{
		No:   0,
		Data: user1,
	})

}

package controller

import "C"
import (
	"encoding/json"
	"github.com/ohko/hst"
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
			No: 200,
			Data: struct {
				msg  string
				data entity.User
			}{msg: "success",
				data: user},
		})
	}

}

//查询所有用户
func GetUserList(c *hst.Context) {
	result, err := service.GetAllUser()
	if err != nil {
		c.JSON2(http.StatusBadRequest, 400, err)
	} else {
		c.JSON2(http.StatusOK, 200, result)
	}
}

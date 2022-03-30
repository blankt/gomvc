package routes

import (
	"git.zx-tech.net/ljhua/hst"
	"simpleMVC/go/controller"
)

func SetRouter() *hst.HST {
	h := hst.New(nil)

	h.POST("/addUser", controller.CreateController)
	h.GET("/getUserList", controller.GetUserList)
	return h
}

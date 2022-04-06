package routes

import (
	"git.zx-tech.net/ljhua/hst"
	"simpleMVC/go/controller"
	"time"
)

func SetRouter() *hst.HST {
	//h := hst.New(nil)
	//
	//h.POST("/addUser", controller.CreateController)
	//h.GET("/getUserList", controller.GetUserList)
	//return h

	h := hst.New(nil)

	//初始化session
	h.SetSession(hst.NewSessionMemory("", "/", "HST_SESSION", time.Minute))

	h.HandleFunc("/addUser", controller.CreateController)
	h.HandleFunc("/getUserList", controller.GetUserList)
	h.RegisterHandle(nil, &controller.User{})
	h.HandleFunc("/login", controller.Login)

	return h
}

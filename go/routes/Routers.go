package routes

import (
	"github.com/ohko/hst"
	"simpleMVC/go/controller"
)

func SetRouter() *hst.HST {
	h := hst.New(nil)

	h.HandleFunc("/addUser", controller.CreateController)
	h.HandleFunc("/getUserList", controller.GetUserList)
	return h
}

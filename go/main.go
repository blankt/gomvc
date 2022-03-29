package main

import (
	"simpleMVC/go/dao"
	"simpleMVC/go/entity"
	"simpleMVC/go/routes"
)

func main() {
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	//绑定模型
	dao.SqlSession.AutoMigrate(&entity.User{})

	r := routes.SetRouter()
	r.ListenHTTP(":8080")
}

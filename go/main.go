package main

import (
	"fmt"
	"simpleMVC/go/dao"
	"simpleMVC/go/entity"
	"simpleMVC/go/routes"
	"time"

	"git.zx-tech.net/ljhua/hst"
)

func main() {
	//初始化数据库连接
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	//绑定模型
	dao.SqlSession.AutoMigrate(&entity.User{})

	//打印sql日志
	dao.SqlSession.LogMode(true)

	r := routes.SetRouter()
	go r.ListenHTTP(":8080")
	//go r.ListenHTTPS(":8081", "server.crt", "server.key")

	//bo := hst.MakeTLSFile("123", "123", "123", "/resources/", "test", "534464210@qq.com")
	//if bo {
	//	log.Printf("make file err")
	//}
	//go r.ListenTLS(":443", "ca.crt", "server.crt", "server.key")

    fmt.Println("测试CI")

	hst.Shutdown(time.Second*5, r)

}

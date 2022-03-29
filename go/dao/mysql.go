package dao

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
)

var SqlSession *gorm.DB

type conf struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
	Port     string `yaml:"post"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("resources/application.yaml")
	if err != nil {
		log.Fatal("读取配置文件失败:", err)
	}

	err1 := yaml.Unmarshal(yamlFile, c)
	if err1 != nil {
		log.Fatal("配置文件转换失败")
	}

	return c
}

func InitMySql() (err error) {
	var c conf
	//获取yaml配置参数
	conf := c.getConf()
	//将yaml配置参数拼接成连接数据库的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
	//连接数据库
	SqlSession, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//验证数据库连接是否成功，若成功，则无异常
	return SqlSession.DB().Ping()
}

func Close() {
	SqlSession.Close()
}

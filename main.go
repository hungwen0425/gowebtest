package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gowebtest/controllers"
	"gowebtest/models"
)

func init() {
	models.RegisterDB()
}

func main() {
	//開發模式
	orm.Debug = true
	//自動建表
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Run()
}

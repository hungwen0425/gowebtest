package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true"

	account := this.Ctx.GetCookie("account")
	pwd := this.Ctx.GetCookie("pwd")

	fmt.Println("Login()方法中：account:", account, ",password:", pwd)

	if isExit {
		this.Ctx.SetCookie("account", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 301)
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	account := this.Input().Get("account")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("account") == account && beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		this.Ctx.SetCookie("account", account, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		//this.SetSecureCookie("account", account, maxAge, "/")
		//this.SetSecureCookie("pwd", pwd, maxAge, "/")
	}

	this.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	check, err := ctx.Request.Cookie("account")
	if err != nil {
		return false
	}

	account := check.Value

	check, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := check.Value
	return beego.AppConfig.String("account") == account && beego.AppConfig.String("pwd") == pwd
}

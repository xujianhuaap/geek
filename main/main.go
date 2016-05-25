package main

import (
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/geek/controller"
)
type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	for _,api:=range controller.Apis {
		beego.Router(api, &MainController{})
	}

	beego.Run()
}

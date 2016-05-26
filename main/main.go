package main
import (
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/geek/controller"
	"github.com/xujianhuaap/geek/models"
	"github.com/xujianhuaap/geek/utils"
	"strings"
)
type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	ctx:=this.Ctx;
	if(strings.EqualFold(ctx.Request.RequestURI,controller.API_USER)){
		user:=models.User{"xujianhua",001,27,true}
		response:=models.Response{Code:200,Message:"成功",Content:user}
		userStr,err:=utils.ToJson(response)
		if(err==nil){
			ctx.WriteString(userStr)
		}else{
			utils.PanicError(err)//错误诊断
		}
	}else{
		ctx.WriteString(ctx.Request.RequestURI)
	}
}

func main() {
	//自动注册路由
	for _,api:=range controller.Apis {
		beego.Router(api, &MainController{})
	}
	beego.Run()
}

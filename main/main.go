package main
import (
	"github.com/xujianhuaap/geek/controller"
	"github.com/xujianhuaap/geek/models"
	"github.com/xujianhuaap/geek/utils"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

var ormer orm.Ormer
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
func (this *MainController)Post() {
	ctx:=this.Ctx;
	if(strings.Contains(ctx.Request.RequestURI,controller.API_USER_REGISTER)){
		user:=models.User{"xujianhua",001,27,true}
		_,err:=ormer.Insert(user)
		var response models.Response
		if(err!=nil){
			utils.PanicError(err)//错误诊断
			response=models.Response{Code:300,Message:"失败",Content:new(models.User)}
		}else{
			response=models.Response{Code:200,Message:"成功",Content:user}
		}

		userStr,err:=utils.ToJson(response)
		if(err==nil){
			ctx.WriteString(userStr)
		}else{
			ctx.WriteString("-----------")
			utils.PanicError(err)//错误诊断
		}
	}else{
		ctx.WriteString(ctx.Request.RequestURI)
	}
}
func main() {
	//数据库初始化
	ormer=orm.NewOrm()//全局执行一次
	ormer.Using("default")
	//自动注册路由
	for _,api:=range controller.Apis {
		beego.Router(api, &MainController{})
	}
	beego.Run()
}

func init()  {
	maxIdle:=30;//最大的闲置数目
	maxConnect:=30;//最大的在线访问
	//orm.RegisterDriver("mysql", orm.DRMySQL) mysql已经默认注册过了
	orm.RegisterDataBase("default","mysql","root:123456@/orm_geek?charset=utf8",maxIdle,maxConnect)
}


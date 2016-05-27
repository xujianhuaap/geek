package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type User struct  {
	Name string `orm:"no null"`
	Userid int `orm:"pk;auto"` //设置为主键
	Age int
	Gender bool //true man;false:womamn
}

func init()  {
	orm.RegisterModel(new(User))
	beego.BeeLogger.Debug("Register init")
}


package controller

const (
	//欢迎界面
	API_WELCOME="/welcome"

	//用户系统
	API_USER="/user"
	API_USER_FAVOURIE_PRODUC="/user/favourite"

	//产品系统
	API_PRODUCT="/product"

	//交易系统
	API_DEAL="/deal"

)
var Apis=[]string{
	//欢迎界面
	API_WELCOME,
	//用户系统
	API_USER,
	API_USER_FAVOURIE_PRODUC,
	//产品系统
	API_PRODUCT,
	//交易系统
	API_DEAL,
}
// http:// 192.168.23.51:9090/welcome
// http://192.168.23.51:9090/user
// http:// 192.168.23.51:9090/user/favourite
// http:// 192.168.23.51:9090/product
// http:// 192.168.23.51:9090/deal


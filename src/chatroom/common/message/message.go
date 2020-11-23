package message

//确定一些消息类型
const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)
type Message struct {
	Type string `json:"type"`//消息类型
	Data string `json:"data"`//消息值
}

//定义两个消息..后面再增加
type LoginMes struct {
	UserId int `json:"userId"`//用户ID
	UserPwd string `json:"userPwd"`//用户密码
	UserName string `json:"userName"`//用户名称
}

type LoginResMes struct {
	Code int `json:"code"`//状态码，500未注册，200成功
	Error string `json:"error"`//返回错误信息
}

type RegisterMes struct {
	//...
}
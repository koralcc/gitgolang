package message

//消息类型常量
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的内容
}

//先定义两个消息，后面需要再增加
type LoginMes struct {
	UserId   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
}

//服务器返回的登陆消息
type LoginResMes struct {
	Code    int    `json:"code"` //返回的状态码 500 表示该用户未注册 200 表示登录成功
	UserIds []int  //增加字段，保存用户id的切片
	Error   string `json:"error"` //返回错误信息
}

type RegisterMes struct {
	User User //类型就是User结构体
}
type RegisterResMes struct {
	Code  int    `json:"code"` //400：表示该用户已经占有，200表示注册成功
	Error string `json:"error"`
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId     int `json:"userId"`
	UserStatus int `json:"status"`
}

//端消息结构体
type SmsMes struct {
	User
	Content string `json:"content"`
}

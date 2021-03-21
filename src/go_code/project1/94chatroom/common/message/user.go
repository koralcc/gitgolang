package message

type User struct {
	//为了序列化和反序列化成功，用户信息的字符串key和结构体的tag保持一致
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"uerName"`
	UserStatus int    `json:"userStatus"`
}

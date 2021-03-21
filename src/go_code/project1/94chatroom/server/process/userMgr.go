package process

import "fmt"

//因为UserMgr 实例在服务器端有且只有一个
//因为在很多地方都会使用到，因此，我们将其定义为全局变量

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UseProcess
}

//完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UseProcess, 1024),
	}
}

//完成对onlineUsers添加

func (this *UserMgr) AddOnlineUser(up *UseProcess) {
	this.onlineUsers[up.UserId] = up
}

func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//获取当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UseProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UseProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok { //说明要查找的这个用户，当前不在线
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}

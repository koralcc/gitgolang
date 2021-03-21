package model

import (
	"encoding/json"
	"fmt"
	"go_code/project1/94chatroom/common/message"

	"github.com/garyburd/redigo/redis"
)

//服务器启动后，就初始化一个UserDao实例，所有的DML操作都公用这个
var (
	MyUserDao *UserDao
)

// 定义一个UserDao结构体
// 完成对User结构体的各种操作

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//在UserDao应该提供哪些方法
//1.根据用户id获取user
func (this *UserDao) getUserById(conn redis.Conn, id int) (user User, err error) {
	//通过给定id去redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { //表示在users哈希表中，没有找到对应id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//注意这里的&user，此时返回参数是结构体，值类型，不需要实例化，反序列化传&user
	//当返回参数类型是指针时，引用类型，要实例化分配地址，在传user即可
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//完成登陆的校验 Login
//1. login完成对用户的校验
//2 .如果用户id和pwd都正确，则返回一个user实例
//3 .如果用户的id或pwd有误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	//先从UserDao的连接池取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//redis进行注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err:", err)
		return
	}
	return

}

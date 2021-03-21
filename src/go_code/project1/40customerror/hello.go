package main

import "errors"

func main() {
	err := readconf("fafsad")
	if err != nil {
		//如果读取文件错误,就输出这个错误,并终止程序
		panic(err)
	}
}

func readconf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误") //定义错误
	}
}

//自定义错误
// Go中,支持自定义错误,使用errors.New()和panic内置函数组合完成
// 1)errors.New("错误说明"),会返回一个error类型的值,标识一个错误
// 2)panic内置函数,接收一个interface{}类型的值(也就是任何值)作为参数,可以接收error类型的变量,输出错误信息,并退出程序

package main

import (
	"fmt"
	"go_code/project1/27package/utils" //gopath配置在环境变量中，全路径 gopath/src+引入路径
)

func main() {
	fmt.Println(utils.Calc(1, 2, "+"))
}

//包
//基本概念：Go的每一个文件都是属于一个包的，也就是说Go是以包的形式来管理文件和项目目录结构的

//包使用细节
// 1)在给一个文件打包时，该包对应一个文件夹，比如这里的utils文件夹对应的包名就是utils,文件的包名通常和文件所在的文件夹名一致，一般为小写字母
// 2)当一个文件要使用其他包函数或变量时，需要先引入对应的包
// (1)引入方式1:import "包名"
// (2)引入方式2:
// import{
// 	"包名"
// 	"包名"
// }
// 3)package 指令在文件第一行，然后是import指令
// 4)在import包时，路径从 $GOPATH的src下开始，编译器会自动从src下开始引入
// 5)同一个包的不同文件，方法和变量直接引用
// 6)为了让其他包的文件可以访问到本包的函数，则该函数名的首字母需要大写，类似其他语言的public，这样才能跨包访问
// 7)在访问其他包函数、变量时，其语法是包名.函数名(变量名)
// 8)如果包名较长，Go支持给包取名，注意细节:取别名后，原来的包名就不能使用了‘
//    9)如果你需要编译成一个可执行程序文件，就需要将这个包声明为main，即package main,这个就是语法规范，如果你是写一个库，包名可以自定义
//    go build -o bin/my.exe go_code/project1/27package/main //指定编译成bin下的my.exe,编译文件夹(不需带上src，编译器自动加上src)，注意，这个main是指装main文件的文件夹
//    go build -o bin/my.exe src/go_code/project1/27package/main/main.go,编译文件需要带上src等全路径

// 关于编译的说明
// 1)编译的指令，在项目目录下，编译路径不需要加src，编译器会自动带上
// 2)编译时需要编译main包所在的文件夹
// 3)项目的目录结构最好按照规范来组织
// 4)编译后生成一个有默认名的可执行文件，在%gopath目录下，可指定名字和目录： go build -o bin/my.exe go_code/project1/27package/main

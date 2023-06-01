// echo 输出命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// s := "" 短变量声明，通常在函数内部使用，不适合包级别的变量
// var s string 默认初始化空字符串的“”

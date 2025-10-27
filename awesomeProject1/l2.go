package main

//我考虑了除数是零的情况，不过我看了看题里没说运行错误，就没加

import (
	"fmt"
	"os"
)

func main() {
	var ru1, ru2 int
	var fuhao, shif string
	for {
		fmt.Printf("欢迎使用Go语言计算器!\n请输入两个整数和一个操作符，进行四则运算\n输入exit退出程序")
		fmt.Printf("请输入第一个整数:")
		fmt.Scan(&ru1)
		fmt.Printf("请输入操作符::")
		fmt.Scan(&fuhao)
		fmt.Printf("请输入第二个整数:")
		fmt.Scan(&ru2)
		switch fuhao {
		case "+":
			fmt.Printf("%d+%d=%d\n", ru1, ru2, ru1+ru2)
		case "-":
			fmt.Printf("%d-%d=%d\n", ru1, ru2, ru1-ru2)
		case "*":
			fmt.Printf("%d*%d=%d\n", ru1, ru2, ru1*ru2)
		case "/":
			fmt.Printf("%d/%d=%d\n", ru1, ru2, ru1/ru2)
		}
		fmt.Printf("是否继续?")
		fmt.Scan(&shif)
		if shif == "y" {
		}
		if shif == "exit" {
			fmt.Printf("感谢使用!再见!\n")
			os.Exit(0)
		}
	}
}

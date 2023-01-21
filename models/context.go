package models

import "fmt"

func PrintDescription() {
	fmt.Println("***************************")
	fmt.Println()
	fmt.Println("1.倒计时关机\t2.定时关机\t3.取消关机计划\t0.退出")
	fmt.Println()
	fmt.Println("***************************")
}

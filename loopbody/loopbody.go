package loopbody

import (
	"ShutDown/cmd"
	"ShutDown/function"
	"ShutDown/models"
	"fmt"
	"os"
)

func LoopBody(key int) {
	var timeT models.Time
	switch key {
	case 1:
		//倒计时关机
		function.Countdown(&timeT)
	case 2:
		//定时关机
		function.ShutdownAt(&timeT)
	case 3:
		//取消关机
		cmd.AbortShutDown()
		os.Exit(0)
	case 0:
		//退出
		os.Exit(0)
	default:
		fmt.Println("输入有误")
	}
}

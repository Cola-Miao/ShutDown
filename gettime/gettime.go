package gettime

import (
	"ShutDown/models"
	"ShutDown/utils"
	"fmt"
)

func GetTime(time *models.Time) {
	var err error
	fmt.Println("输入时")
	_, err = fmt.Scanln(&time.Hour)
	utils.RecordError(err)
	fmt.Println("输入分")
	_, err = fmt.Scanln(&time.Minute)
	utils.RecordError(err)
	fmt.Println("输入秒")
	_, err = fmt.Scanln(&time.Second)
	utils.RecordError(err)
}

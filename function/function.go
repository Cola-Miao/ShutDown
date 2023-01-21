package function

import (
	"ShutDown/cmd"
	"ShutDown/gettime"
	"ShutDown/models"
	"os"
	"strconv"
	"time"
)

func Countdown(timeT *models.Time) {
	gettime.GetTime(timeT)
	countdown := timeT.Hour*60*60 + timeT.Minute*60 + timeT.Second
	countdownStr := strconv.Itoa(countdown)
	cmd.ShutDown(countdownStr)
	os.Exit(0)
}

func ShutdownAt(timeT *models.Time) {
	gettime.GetTime(timeT)
	now := time.Now()
	countdown := (timeT.Hour-now.Hour())*60*60 + (timeT.Minute-now.Minute())*60 + (timeT.Second - now.Second())
	countdownStr := strconv.Itoa(countdown)
	cmd.ShutDown(countdownStr)
	os.Exit(0)
}

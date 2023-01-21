package cmd

import (
	"ShutDown/utils"
	"os/exec"
)

func ShutDown(countdown string) {
	cmd := exec.Command("shutdown", "-s", "-t", countdown)
	err := cmd.Run()
	utils.RecordError(err)
}

func AbortShutDown() {
	cmd := exec.Command("shutdown", "-a")
	err := cmd.Run()
	utils.RecordError(err)
}

package main

import (
	"ShutDown/loopbody"
	"ShutDown/models"
	"ShutDown/utils"
)

func main() {
	models.PrintDescription()
	key := utils.GetKey()

	for {
		loopbody.LoopBody(key)
	}
}

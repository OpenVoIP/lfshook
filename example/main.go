package main

import (
	"github.com/OpenVoIP/lfshook"
)

func main() {
	log := lfshook.NewFileLogger("/var/log/info.log", "/var/log/error.log", 1)

	log.Info("hello info")
	log.Error("hello error")
}

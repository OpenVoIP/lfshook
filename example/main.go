package main

import (
	"github.com/OpenVoIP/lfshook"
)

func main() {
	log := lfshook.NewFileLogger("/var/log/info.log", "/var/log/error.log")
	log.Info("info hello")
	log.Info("info hello")
	log.Error("error hello")
}

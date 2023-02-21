package main

import (
	"time"
)

func logger(input string) {
	fileMkDir("./log")
	println("[" + time.Now().Format("2006-02-01 15:04:05") + "] " + input)
	fileAddLine("["+time.Now().Format("2006-02-01 15:04:05")+"] "+input, "./log/"+time.Now().Format("2006-02-01")+".log")
}

func errorPanic(err error, logBefore ...string) {
	if err != nil {
		for _, log := range logBefore {
			logger(log)
		}
		logger(err.Error())
		panic("\n----------------ERROR----------------\n")
	}
}

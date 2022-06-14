package main

import (
	"easyLog"
	"fmt"
	"time"
)

func consoleTest() {
	log, err := easyLog.NewConsoleLogger("error")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		log.Debug("这是一条Debug日志")
		log.Trace("这是一条Trace日志")
		log.Info("这是一条Info日志")
		log.Warn("这是一条Warn日志")
		id := 123
		log.Error("这是一条Error日志,id:%d\n", id)
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 3)
	}
}
func fileTest() {
	log, err := easyLog.NewFileLogger("debug", "./easyLog_demo/", "test.log", 10*1024)
	//fmt.Printf("fileLogger对象创建成功：%v\n", log)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		log.Debug("这是一条Debug日志")
		log.Trace("这是一条Trace日志")
		log.Info("这是一条Info日志")
		log.Warn("这是一条Warn日志")
		id := 123
		log.Error("这是一条Error日志,id:%d\n", id)
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 3)
	}
}
func main() {
	//consoleTest()
	fileTest()
}

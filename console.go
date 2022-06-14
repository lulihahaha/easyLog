package easyLog

import (
	"fmt"
	"time"
)

// ConsoleLogger 日志类
type ConsoleLogger struct {
	Level LogLevel
}

// logPrint 打印日志记录
func (c *ConsoleLogger) logPrint(lv LogLevel, format string, a ...interface{}) {
	if c.Level <= lv {
		level, err := stringToLevel(lv)
		if err != nil {
			fmt.Println(err)
			return
		}
		funcName, fileName, lineNo := getInfo(3)
		t := time.Now().Format("2006-01-02 15:04:05.000")
		msg := fmt.Sprintf(format, a...)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", t, level, fileName, funcName, lineNo, msg)
	}
}

// NewConsoleLogger ConsoleLogger类的构造函数
func NewConsoleLogger(msg string) (*ConsoleLogger, error) {
	level, err := parseLevel(msg)
	if err != nil {
		return nil, err
	}
	return &ConsoleLogger{
		Level: level,
	}, nil
}

// Debug Debug级别的日志信息
func (c *ConsoleLogger) Debug(format string, a ...interface{}) {
	c.logPrint(DEBUG, format, a...)
}

// Trace Trace级别的日志信息
func (c *ConsoleLogger) Trace(format string, a ...interface{}) {
	c.logPrint(TRACE, format, a...)
}

// Info Info级别的日志信息
func (c *ConsoleLogger) Info(format string, a ...interface{}) {
	c.logPrint(INFO, format, a...)
}

// Warn Warn级别的日志信息
func (c *ConsoleLogger) Warn(format string, a ...interface{}) {
	c.logPrint(WARN, format, a...)
}

// Error Error级别的日志信息
func (c *ConsoleLogger) Error(format string, a ...interface{}) {
	c.logPrint(ERROR, format, a...)
}

// Fatal Fatal级别的日志信息
func (c *ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.logPrint(FATAL, format, a...)
}

package easyLog

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

/*
日志库：
1.日志级别：Debug Trace Info Warn Error Fatal
2.可以往文件中写入，也可以往终端输出,可拓展
3.支持开发/生产环境切换，开发模式支持全级别，生产模式支持Info以上的级别
4.日志记录要有时间、行号、文件名、日志信息、日志级别等
5.日志文件要切割
*/

// LogLevel 定义日志的级别，包括Debug Trace Info Warn Error Fatal
type LogLevel uint32

const (
	UNKOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARN
	ERROR
	FATAL
)

// EasyLogger EasyLogger接口
type EasyLogger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

// parseLevel 将字符串解析成对应的LogLevel
func parseLevel(msg string) (LogLevel, error) {
	s := strings.ToLower(msg)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warn":
		return WARN, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		{
			err := errors.New("日志级别错误，log level error")
			return UNKOWN, err
		}

	}
}

// stringToLevel 将字符串解析成对应的LogLevel
func stringToLevel(lv LogLevel) (s string, err error) {
	switch lv {
	case DEBUG:
		return "Debug", nil
	case TRACE:
		return "Trace", nil
	case INFO:
		return "Info", nil
	case WARN:
		return "Warn", nil
	case ERROR:
		return "Error", nil
	case FATAL:
		return "Fatal", nil
	default:
		{
			err = errors.New("stringToLevel( failed!")
			return "Unkown", err
		}
	}
}

// getInfo 获得函数名，文件名，行号等相关信息
func getInfo(skip int) (funcName, file string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	return
}

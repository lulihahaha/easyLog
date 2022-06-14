package easyLog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 文件日志类
type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
}

// NewFileLogger 创建文件日志类对象
func NewFileLogger(levelStr, fp, fn string, maxSize int64) (*FileLogger, error) {
	logLevel, err := parseLevel(levelStr)
	if err != nil {
		fmt.Println("parseLevel() failed")
		return nil, err
	}
	f1 := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1, nil
}

// initFile 初始化文件句柄
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open errorLog file failed,err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

// checkSize 判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

// splitFile 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v", err)
		return nil, err
	}
	file.Close()
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	os.Rename(logName, newLogName)
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}
	return fileObj, nil
}

// logPrint 往文件中打印日志
func (f *FileLogger) logPrint(lv LogLevel, format string, a ...interface{}) {
	if f.Level <= lv {
		level, err := stringToLevel(lv)
		if err != nil {
			fmt.Println(err)
			return
		}
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		t := time.Now().Format("2006-01-02 15:04:05.000")
		msg := fmt.Sprintf(format, a...)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", t, level, fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", t, level, fileName, funcName, lineNo, msg)
		}
	}
}

// close 关闭文件
func (f *FileLogger) close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

// Debug Debug级别的日志信息
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.logPrint(DEBUG, format, a...)
}

// Trace Trace级别的日志信息
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.logPrint(TRACE, format, a...)
}

// Info Info级别的日志信息
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.logPrint(INFO, format, a...)
}

// Warn Warn级别的日志信息
func (f *FileLogger) Warn(format string, a ...interface{}) {
	f.logPrint(WARN, format, a...)
}

// Error Error级别的日志信息
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.logPrint(ERROR, format, a...)
}

// Fatal Fatal级别的日志信息
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.logPrint(FATAL, format, a...)
}

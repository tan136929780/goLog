/*
@Time : 2020/12/14 下午3:43
@Author : tan
@File : PrintLog
@Software: GoLand
*/
package logger

import (
	"fmt"
	"path"
	"runtime"
)

type PrintLogger struct {
	logLevel       int    `json:"log_level"`
	logLevelString string `json:"log_level_string"`
	message        string `json:"message"`
}

/*
Info level log
*/
func (printLogger *PrintLogger) Info(message string, fileName string, args ...interface{}) {
	if LEVEL_INFO > printLogger.logLevel {
		return
	}
	printLogger.logLevelString = levelString[LEVEL_INFO]
	printLogger.message = message
	processPrintLog(printLogger)
}

/*
Trace level log
*/
func (printLogger *PrintLogger) Trace(message string, fileName string, args ...interface{}) {
	if LEVEL_TRACE > printLogger.logLevel {
		return
	}
	printLogger.logLevelString = levelString[LEVEL_TRACE]
	printLogger.message = message
	processPrintLog(printLogger)
}

/*
Debug level log
*/
func (printLogger *PrintLogger) Debug(message string, fileName string, args ...interface{}) {
	if LEVEL_DEBUG > printLogger.logLevel {
		return
	}
	printLogger.logLevelString = levelString[LEVEL_DEBUG]
	printLogger.message = message
	processPrintLog(printLogger)

}

/*
Error level log
*/
func (printLogger *PrintLogger) Error(message string, fileName string, args ...interface{}) {
	if LEVEL_ERROR > printLogger.logLevel {
		return
	}
	printLogger.logLevelString = levelString[LEVEL_ERROR]
	printLogger.message = message
	processPrintLog(printLogger)
}

/*
Fatal level log
*/
func (printLogger *PrintLogger) Fatal(message string, fileName string, args ...interface{}) {
	if LEVEL_FATAL > printLogger.logLevel {
		return
	}
	printLogger.logLevelString = levelString[LEVEL_FATAL]
	printLogger.message = message
	processPrintLog(printLogger)
}

/*
errorLog error log
*/
func (logger *PrintLogger) ErrorLog(message, fileName string, args ...interface{}) {
	logger.Error(message, path.Join("errorLog", fileName))
}

/*
setConfig set params
*/
func (printLogger *PrintLogger) setConfig(size int, count int, args ...interface{}) {
	//
}

/*
processPrintLog print log
*/
func processPrintLog(printLogger *PrintLogger) {
	pc, fileNamePath, line, _ := runtime.Caller(2)
	pcName := runtime.FuncForPC(pc).Name()
	fmt.Printf("Level: %v\n", printLogger.logLevelString)
	fmt.Printf("File: %v\n", fileNamePath)
	fmt.Printf("Func: %v\n", pcName)
	fmt.Printf("Line: %v\n", line)
	fmt.Printf("Message: %v\n", printLogger.message)
}

/*
@Time : 2020/12/14 下午3:39
@Author : tan
@File : fileLog
@Software: GoLand
*/
package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//defect log dir
var (
	ProjectDir, _ = os.Getwd()
	LogDir        = path.Join(ProjectDir, "log")
	ErrorLogDir   = path.Join(ProjectDir, "errorLog")
	LogName       = "app"
	LogSuffix     = ".log"
)

//FileLogger logger struct
type FileLogger struct {
	logLevel        int
	logSize         int
	logMaxFileCount int
	message         string
	fileName        string
}

/*
Info level log
*/
func (fileLogger *FileLogger) Info(message string, fileName string, args ...interface{}) {
	if LEVEL_INFO > fileLogger.logLevel {
		return
	}
	fileLogger.message = message
	fileLogger.fileName = fileName
	processFileLog(fileLogger, levelString[LEVEL_INFO])
}

/*
Trace level log
*/
func (fileLogger *FileLogger) Trace(message string, fileName string, args ...interface{}) {
	if LEVEL_TRACE > fileLogger.logLevel {
		return
	}
	fileLogger.message = message
	fileLogger.fileName = fileName
	processFileLog(fileLogger, levelString[LEVEL_TRACE])
}

/*
Debug level log
*/
func (fileLogger *FileLogger) Debug(message string, fileName string, args ...interface{}) {
	if LEVEL_DEBUG > fileLogger.logLevel {
		return
	}
	fileLogger.message = message
	fileLogger.fileName = fileName
	processFileLog(fileLogger, levelString[LEVEL_DEBUG])
}

/*
Error level log
*/
func (fileLogger *FileLogger) Error(message string, fileName string, args ...interface{}) {
	if LEVEL_ERROR > fileLogger.logLevel {
		return
	}
	fileLogger.message = message
	fileLogger.fileName = fileName
	processFileLog(fileLogger, levelString[LEVEL_ERROR])
}

/*
Fatal level log
*/
func (fileLogger *FileLogger) Fatal(message string, fileName string, args ...interface{}) {
	if LEVEL_FATAL > fileLogger.logLevel {
		return
	}
	fileLogger.message = message
	fileLogger.fileName = fileName
	processFileLog(fileLogger, levelString[LEVEL_FATAL])
}

/*
errorLog error log
*/
func (logger *FileLogger) ErrorLog(message, fileName string, args ...interface{}) {
	logger.Error(message, path.Join("errorLog", fileName))
}

/*
setConfig set params
*/
func (fileLogger *FileLogger) setConfig(size, count int, args ...interface{}) {
	fileLogger.logSize = size
	fileLogger.logMaxFileCount = count
}

/*
processFileLog create log process
*/
func processFileLog(fileLogger *FileLogger, level string) {
	//format file info
	logFileToWrite, logDir := formatFileNameAndDir(fileLogger)
	//check dir and create if not exist
	checkDirIfNotExistCreate(logDir)
	//check max file count
	checkMaxFile(fileLogger, logDir)
	//format log
	log := createLog(fileLogger, level)
	//doFileAction
	fileWrite := doFileAction(fileLogger, logFileToWrite)
	defer fileWrite.Close()
	//write log file
	fmt.Fprint(fileWrite, log)
}

/*
formatFileNameAndDir get dir
*/
func formatFileNameAndDir(fileLogger *FileLogger) (logFileToWrite, logDir string) {
	var logFile string
	if fileLogger.fileName == "" {
		logFile = strings.Replace(fileLogger.fileName, "", LogName, 1)
	} else {
		logFile = fileLogger.fileName
	}
	logFileToWrite = path.Join(LogDir, logFile, logFile+LogSuffix)
	logDir = path.Join(LogDir, logFile)
	return logFileToWrite, logDir
}

/*
check dir is exist
*/
func checkDirIfNotExistCreate(dir string) bool {
	_, error := os.Stat(dir)
	if !os.IsNotExist(error) {
		return false
	}
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return false
	}
	return true
}

/*
check whether the ole file need cut size
*/
func checkCutFile(fileLogger *FileLogger, file *os.File) bool {
	stat, _ := file.Stat()
	return int(stat.Size()) > fileLogger.logSize
}

/*
delete files if file count is overflow
*/
func checkMaxFile(fileLogger *FileLogger, dir string) bool {
	rd, error := ioutil.ReadDir(dir)
	if error != nil {
		return false
	}
	count := len(rd)
	rdSlice := rd[:]
	if count >= fileLogger.logMaxFileCount {
		for i := 0; i < len(rdSlice)-fileLogger.logMaxFileCount; i++ {
			file := rdSlice[i+1]
			os.Remove(path.Join(dir, file.Name()))
		}
	}
	return true
}

/*
createLog create log
*/
func createLog(fileLogger *FileLogger, level string) string {
	pc, fileNamePath, line, _ := runtime.Caller(3)
	pcName := runtime.FuncForPC(pc).Name()
	log := fmt.Sprintf("{\n    file: %v\n    func: %v\n	line: %v\n    level: %v\n    time: %v\n    message: %v\n}\n", fileNamePath, pcName, strconv.Itoa(line), level, time.Now().Format("2006-01-02 15:04:05"), fileLogger.message)
	return log
}

/*
doFileAction do file action
*/
func doFileAction(fileLogger *FileLogger, logFileToWrite string) *os.File {
	pathName := path.Dir(logFileToWrite)
	_, error := os.Stat(pathName)
	if os.IsNotExist(error) {
		os.MkdirAll(pathName, 0755)
	}
	fileWrite, error := os.OpenFile(logFileToWrite, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	ifCut := checkCutFile(fileLogger, fileWrite)
	if ifCut {
		fileWrite = renameFile(fileWrite)
	}

	if error != nil {
		defer fileWrite.Close()
		return fileWrite
	}
	return fileWrite
}

/*
renameFile create new file when file size is overflow
*/
func renameFile(read *os.File) (write *os.File) {
	defer read.Close()
	timeNow := time.Now().UnixNano()
	oldFile := read.Name()
	newFile := fmt.Sprintf("%v.%v", oldFile, strconv.Itoa(int(timeNow)))
	err := os.Rename(oldFile, newFile)
	if err != nil {
		return read
	}
	write, writeErr := os.OpenFile(oldFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if writeErr != nil {
		defer write.Close()
		return write
	}
	return write
}

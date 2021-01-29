/*
@Time : 2020/12/14 下午3:34
@Author : tan
@File : log
@Software: GoLand
*/
package logger

//LEVEL
const (
	LEVEL_NOMAL = iota
	LEVEL_INFO
	LEVEL_TRACE
	LEVEL_DEBUG
	LEVEL_ERROR
	LEVEL_FATAL
)

//LOG TYPE
const (
	MODE_DEFAULT_LOGGER = iota
	MODE_FILE_LOGGER
	MODE_PRINT_LOGGER
)

//LEVEL STRING
var levelString = []string{
	"NOMAL",
	"INFO",
	"TRACE",
	"DEBUG",
	"ERROR",
	"FATAL",
}

/*
Logger interface
*/
type Logger interface {
	Info(string, string, ...interface{})
	Trace(string, string, ...interface{})
	Debug(string, string, ...interface{})
	Error(string, string, ...interface{})
	Fatal(string, string, ...interface{})
	ErrorLog(string, string, ...interface{})
	setConfig(int, int, ...interface{})
}

/*
NewLogger set log params

log1 := *logger.NewLogger(logger.MODE_DEFAULT_LOGGER, logger.LEVEL_FATAL)

log2 := *logger.NewLogger(logger.MODE_FILE_LOGGER, logger.LEVEL_FATAL)

log3 := *logger.NewLogger(logger.MODE_FILE_LOGGER, logger.LEVEL_FATAL)

log4 := *logger.NewLogger(logger.MODE_PRINT_LOGGER, logger.LEVEL_FATAL)

log1.Debug("1234567890", "")

log1.Fatal("1234567890", "test")

log2.Info("1234567890", "")

log3.Trace("1234567890", "")

log4.Error("1234567890", "")
*/
func NewLogger(logMode int, level int) *Logger {
	var logger Logger
	switch logMode {
	case MODE_FILE_LOGGER:
		logger = &FileLogger{
			logLevel:        level,
			message:         "",
			fileName:        "",
			logSize:         (20 * 1024 * 1024),
			logMaxFileCount: 5,
		}
		break
	case MODE_PRINT_LOGGER:
		logger = &PrintLogger{
			logLevel:       level,
			logLevelString: "",
			message:        "",
		}
		break
	default:
		break
	}
	return &logger
}

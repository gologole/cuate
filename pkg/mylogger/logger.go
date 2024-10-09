package log

import (
	"cmd/main.go/config"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var MyLogger *logrus.Logger

func Init() {
	MyLogger = logrus.New()

	if config.IsDebug == "TRUE" {
		consoleOutput := logrus.New()
		consoleOutput.SetFormatter(&logrus.TextFormatter{ForceColors: true})
		MyLogger.SetOutput(io.MultiWriter(consoleOutput.Writer(), os.Stdout))
	} else {
		now := time.Now()
		filelog := os.Getenv("LOGS_DIR") + "/" + now.Format("20060102") + ".log"

		file, err := os.OpenFile(filelog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		MyLogger.SetOutput(file)

		MyLogger.SetFormatter(&logrus.JSONFormatter{})
	}

	if config.IsDebug == "TRUE" {

		now := time.Now()
		filelog := "./logs" + now.Format("20060102") + ".log"

		file, err := os.OpenFile(filelog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		MyLogger.SetOutput(io.MultiWriter(os.Stdout, file))
		MyLogger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	} else {
		now := time.Now()
		filelog := os.Getenv("LOGS_DIR") + "/" + now.Format("20060102") + ".log"

		file, err := os.OpenFile(filelog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		MyLogger.SetOutput(file)
		MyLogger.SetFormatter(&logrus.JSONFormatter{})
	}
}

func Trace(args ...interface{}) {
	MyLogger.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	MyLogger.Tracef(format, args...)
}

func Debug(args ...interface{}) {
	MyLogger.Trace(args...)
}

func Debugf(format string, args ...interface{}) {
	MyLogger.Tracef(format, args...)
}

func Info(args ...interface{}) {
	MyLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	MyLogger.Infof(format, args...)
}

func Warning(args ...interface{}) {
	MyLogger.Warning(args...)
}

func Warningf(format string, args ...interface{}) {
	MyLogger.Warningf(format, args...)
}

func Error(args ...interface{}) {
	MyLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	MyLogger.Errorf(format, args...)
}

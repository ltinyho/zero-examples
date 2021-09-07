package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var logFilePath string
var logFileName = "app.log"
var log *logrus.Entry
func AddDirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
func main() {
	logger:=LogInit()
	log = logger.WithField("service","cloud-change")
	for i := 0; i < 10000; i++ {
		log.Info("ok")
		log.Error("ok")
		log.Debug("ok")
	}
}

func LogInit() *logrus.Logger{
	if logFilePath == "" {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		logFilePath = path + "/logs/"
	}
	err := AddDirIfNotExist(logFilePath)
	if err != nil {
		return nil
	} //nolint:errcheck
	fileName := logFilePath + logFileName
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"}) // 设置日志格式

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log", // 分割后的文件名称
		//rotatelogs.WithLinkName(logFileName),      // 生成软链，指向最新日志文件
		rotatelogs.WithRotationSize(1024*1024),
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 设置最大保存时间(7天)
		rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔(1天)
	)
	if err != nil {
		panic(err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))
	return logger
}

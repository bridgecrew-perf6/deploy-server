// Golang program to illustrate the usage of
// fmt.Fprintln() function

// Including the main package
package main

// Importing fmt and os
import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func sss() int {
	fmt.Println("gfghj")
	return 10
}

//var log = log.New()

// Calling main
//func main() {
//	//customFormatter := new(logrus.TextFormatter)
//	//customFormatter.TimestampFormat = "2006-01-02 15:04:05"
//	//logrus.SetFormatter(customFormatter)
//	//customFormatter.FullTimestamp = true
//	//logrus.Info("Hello Walrus after FullTimestamp=true")
//	log.SetFormatter(&log.TextFormatter{
//		DisableColors: false,
//		FullTimestamp: true,
//		TimestampFormat: "2006-01-02 15:04:05",
//	})
//	log.SetReportCaller(true)
//	log.Info("gfhvjhbkjnlksrhgrsh")
//
//	logrus := log.New()
//	//file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//	//if err == nil {
//	//	logrus.Out = file
//	//} else {
//	//log.Info("Failed to log to file, using default stderr")
//	//}
//	logrus.SetFormatter(&log.TextFormatter{
//		DisableColors: false,
//		FullTimestamp: true,
//		TimestampFormat: "2006-01-02 15:04:05",
//	})
//	logrus.SetReportCaller(true)
//	logrus.Info("rwe.hreherherh")
//	//logrus.WithFields(log.Fields{}).Info("A group of walrus emerges from the ocean")
//}

// Create a new instance of the logger. You can have any number of instances.
//var log = logrus.New()

func main() {
	if file, err := os.OpenFile("log1.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666); err == nil {
		log.SetOutput(file)
	}

	// Creates a new logger
	var logTwo = logrus.New()
	if file, err := os.OpenFile("log2.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666); err == nil {
		logTwo.SetOutput(file)
	}

	logTwo.SetFormatter(
		&logrus.TextFormatter{
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)
	logTwo.SetReportCaller(true)
	logTwo.Info("rwe.hreherherh")
	//logrus.WithFields(log.Fields{}).Info("A group of walrus emerges from the ocean")
}

//// config logrus log to local filesystem, with file rotation
//func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
//	baseLogPaht := path.Join(logPath, logFileName)
//	writer, err := rotatelogs.New(
//		baseLogPaht+".%Y%m%d%H%M",
//		rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件
//		rotatelogs.WithMaxAge(maxAge), // 文件最大保存时间
//		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
//	)
//	if err != nil {
//		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
//	}
//	lfHook := lfshook.NewHook(lfshook.WriterMap{
//		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
//		log.InfoLevel:  writer,
//		log.WarnLevel:  writer,
//		log.ErrorLevel: writer,
//		log.FatalLevel: writer,
//		log.PanicLevel: writer,
//	})
//	log.AddHook(lfHook)
//}
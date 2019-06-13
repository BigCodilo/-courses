package logger

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func SetLoggers() {
	fileInfo, _ := os.OpenFile("logger/infoLogger.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	Info = log.New(fileInfo, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	fileError, _ := os.OpenFile("logger/infoError.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	Error = log.New(fileError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

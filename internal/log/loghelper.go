package logMsg

import (
	"log"
	"os"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

func InitializeLoggers() {
	Debug = log.New(os.Stdout, "DEBUG", log.Ldate|log.Ltime|log.Lmicroseconds)
	Info = log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Lmicroseconds)
	Error = log.New(os.Stdout, "ERROR", log.Ldate|log.Ltime|log.Lmicroseconds)
}

func DebugLog(msg string) {
	Debug.Println(msg)
}
func InfoLog(msg string) {
	Info.Println(msg)
}

func ErrorLog(msg string) {
	Error.Println(msg)
}

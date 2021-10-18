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

func WriteDebugLog(msg string) {
	Debug.Println(msg)
}
func WriteInfoLog(msg string) {
	Info.Println(msg)
}

func WriteErrorLog(msg string) {
	Error.Println(msg)
}

package logger

import (
	"log"
	"os"
)

type Logger struct {
	logLogger *log.Logger
}

func New() *Logger {
	return &Logger{
		logLogger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (lg *Logger) Println(v ...interface{}) {
	lg.logLogger.Println(v...)
}

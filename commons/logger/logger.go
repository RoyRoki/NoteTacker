package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

// InitLogger initializes the logger
func InitLogger() {
	Logger = log.New(os.Stdout, "[NOTE] ", log.Ltime)
}

// Info logs informational messages.
func Info(v ...interface{}) {
	Logger.Println(("[INFO]"), v)
}

// Error logs error messages.
func Error(v ...interface{}) {
	Logger.Println("[ERROR]", v)
}

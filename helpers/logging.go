package helpers

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/ramadhanalfarisi/go-codebase/config"
)

var (
	date          = time.Now().Format("2006-01-02")
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	DebugLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(msg string) {
	InfoLogger.Println(msg)
}

func Warning(msg string) {
	WarningLogger.Println(msg)
}

func Debug(msg string) {
	if config.DEBUG == "true" {
		DebugLogger.Println(msg)
	}
}

func Error(err error) {
	var logs string
	pc, fn, line, _ := runtime.Caller(1)
	// Include function name if debugging
	if config.DEBUG == "true" {
		logs = fmt.Sprintf("%s [%s:%s:%d]", err, runtime.FuncForPC(pc).Name(), fn, line)
	} else {
		logs = fmt.Sprintf("%s [%s:%d]", err, fn, line)
	}
	ErrorLogger.Println(logs)
}

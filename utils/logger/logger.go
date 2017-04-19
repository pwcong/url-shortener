package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	Init "github.com/pwcong/url-shortener/init"
)

// LOGFORMAT decide the format of log.
const LOGFORMAT = "%s - %s - %s"

// Log level
const (
	ACCESS = iota
	SERVER
	ERROR
)

var isProd bool

var accessLogger *log.Logger
var errorLogger *log.Logger
var serverLogger *log.Logger

func Log2Access(head string, msg string, explain string) {
	Log(ACCESS, head, msg, explain)
}
func Log2Error(head string, msg string, explain string) {
	Log(ERROR, head, msg, explain)
}
func Log2Server(head string, msg string, explain string) {
	Log(SERVER, head, msg, explain)
}

// Log can print info on console or file that decided by config.
func Log(level int, head string, msg string, explain string) {

	switch level {
	case ACCESS:
		logByLogMode(accessLogger, head, msg, explain)
	case SERVER:
		logByLogMode(serverLogger, head, msg, explain)
	case ERROR:
		logByLogMode(errorLogger, head, msg, explain)

	default:
		log.Printf(LOGFORMAT, head, msg, explain)
	}

}

func logByLogMode(logger *log.Logger, head string, msg string, explain string) {

	if isProd {

		logger.Printf(LOGFORMAT, head, msg, explain)

	} else {

		log.Printf(LOGFORMAT, head, msg, explain)

	}

}

func initLogger() {

	loggerDir := filepath.Join(filepath.Dir(os.Args[0]), "log")

	if _, err := os.Stat(loggerDir); err != nil {
		err := os.MkdirAll(loggerDir, 0666)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	errorLoggerPath := filepath.Join(loggerDir, "error.log")
	f, err := os.OpenFile(errorLoggerPath, os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	errorLogger = log.New(f, "error: ", log.Ldate|log.Ltime|log.Lmicroseconds)

	serverLoggerPath := filepath.Join(loggerDir, "server.log")
	f, err = os.OpenFile(serverLoggerPath, os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	serverLogger = log.New(f, "server: ", log.Ldate|log.Ltime|log.Lmicroseconds)

	accessLoggerPath := filepath.Join(loggerDir, "access.log")
	f, err = os.OpenFile(accessLoggerPath, os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	accessLogger = log.New(f, "access: ", log.Ldate|log.Ltime|log.Lmicroseconds)
}

func initLoggerByLogMode() {

	if Init.Config.Mode == "prod" {

		isProd = true

		initLogger()

	} else {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	}

}

func init() {

	initLoggerByLogMode()
}

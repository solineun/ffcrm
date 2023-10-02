package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

type Logger struct {
	infoLog *log.Logger
	errLog *log.Logger
}

func NewLogger(infoLog, errLog *log.Logger) *Logger {
	if infoLog == nil {
		return &Logger{
			infoLog: DefaultInfo,
			errLog: errLog,
		}
	} else if errLog == nil {
		return &Logger{
			infoLog: infoLog,
			errLog: DefaultErr,
		}
	} else if infoLog == nil || errLog == nil{
		return &Logger{
			infoLog: DefaultInfo,
			errLog: DefaultErr,
		}
	}
	return &Logger{
		infoLog: infoLog,
		errLog: errLog,
	}
}

var DefaultInfo = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var DefaultErr = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func (l *Logger) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	l.errLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (l *Logger) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (l *Logger) Printf(format string, v any) {
	l.infoLog.Printf(format, v)
}

func (l *Logger) Fatal(err error) {
	log.Fatal(err)
}
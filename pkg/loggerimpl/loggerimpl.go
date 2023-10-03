package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

type Logger struct {
	InfoLog *log.Logger
	ErrLog *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
} 

func (l *Logger) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	l.ErrLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (l *Logger) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (l *Logger) Printf(format string, v any) {
	l.InfoLog.Printf(format, v)
}

func (l *Logger) Fatal(err error) {
	log.Fatal(err)
}
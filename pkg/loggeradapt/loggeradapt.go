package loggeradapt

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

type LoggerAdapter struct {
	InfoLog *log.Logger
	ErrLog *log.Logger
}

func NewLoggerAdapter() LoggerAdapter {
	return LoggerAdapter{
		InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
} 

func (l LoggerAdapter) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	l.ErrLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (l LoggerAdapter) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (l LoggerAdapter) Printf(format string, v any) {
	l.InfoLog.Printf(format, v)
}

func (l LoggerAdapter) Fatal(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	l.ErrLog.Output(2, trace)
}
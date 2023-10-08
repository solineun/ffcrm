package logic

import (
	"errors"
	"net/http"

	"github.com/solineun/ffcrm/internal/logic/handler"
	"github.com/solineun/ffcrm/internal/logic/logger"
	"github.com/solineun/ffcrm/internal/logic/storage"
)

type Logic interface {
	logger.Logger
	handler.Handler
}

type LogicAdapter struct {
	errLog 	logger.ErrLogger
	infoLog logger.InfoLogger
	db  storage.FFcrmDB
}

// Home implements Handler.
func (la *LogicAdapter) Home(w http.ResponseWriter, r *http.Request) {
	orders, err := la.db.LatestFiveOrders()
	if err != nil && !errors.Is(err, storage.ErrNoRecord) {
		la.errLog.ServerError(w, err)
	}
	resp := []byte{}
	for _, o := range orders {
		resp = append(resp, []byte(o.Format())...)
	}
	w.Write(resp)
}

//ServerError implements ErrLogger.
func (la *LogicAdapter) ServerError(w http.ResponseWriter, err error) {
	la.errLog.ServerError(w, err)
}

// Fatal implements ErrLogger.
func (la *LogicAdapter) Fatal(err error) {
	la.errLog.Fatal(err)
}

// Printf implements InfoLogger.
func (la *LogicAdapter) Printf(format string, v any) {
	la.infoLog.Printf(format, v)
}

func NewLogicAdapter(errLog logger.ErrLogger, infoLog logger.InfoLogger, db storage.FFcrmDB) *LogicAdapter {
	return &LogicAdapter{
		errLog: errLog,
		infoLog: infoLog,
		db:  db,
	}
}
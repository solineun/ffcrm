package logic

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/solineun/ffcrm/internal/logic/handler"
	"github.com/solineun/ffcrm/internal/logic/logger"
	"github.com/solineun/ffcrm/internal/logic/storage"
	"github.com/solineun/ffcrm/pkg/logadapt"
	"github.com/solineun/ffcrm/pkg/models"
	"github.com/solineun/ffcrm/pkg/models/pg"
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
	if err != nil && !errors.Is(err, models.ErrNoRecord) {
		la.errLog.ServerError(w, err)
	}
	resp := []byte{}
	for _, o := range orders {
		resp = append(resp, []byte(o.Format())...)
	}
	w.Write(resp)
}

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

func NewLogicAdapter(db *sql.DB) *LogicAdapter {
	return &LogicAdapter{
		errLog: logadapt.NewLoggerAdapter(),
		infoLog: logadapt.NewLoggerAdapter(),
		db:  pg.NewFFcrmDB(db),
	}
}
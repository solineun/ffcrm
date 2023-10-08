package logicadapt

import (
	"errors"
	"net/http"

	"github.com/solineun/ffcrm/internal/storage"
	la "github.com/solineun/ffcrm/pkg/loggeradapt"
)

type LogicAdapter struct {
	log la.LoggerAdapter
	db  storage.FFcrmDB
	
}

// Home implements Handler.
func (la *LogicAdapter) Home(w http.ResponseWriter, r *http.Request) {
	orders, err := la.db.LatestFiveOrders()
	if err != nil && !errors.Is(err, storage.ErrNoRecord) {
		la.log.ServerError(w, err)
	}
	resp := []byte{}
	for _, o := range orders {
		resp = append(resp, []byte(o.Format())...)
	}
	w.Write(resp)
}

//ServerError implements ErrLogger.
func (la *LogicAdapter) ServerError(w http.ResponseWriter, err error) {
	la.log.ServerError(w, err)
}

// Fatal implements ErrLogger.
func (la *LogicAdapter) Fatal(err error) {
	la.log.Fatal(err)
}

// Printf implements InfoLogger.
func (la *LogicAdapter) Printf(format string, v any) {
	la.log.Printf(format, v)
}

func NewLogicAdapter(db storage.FFcrmDB) *LogicAdapter {
	return &LogicAdapter{
		log: la.NewLoggerAdapter(),
		db:  db,
	}
}
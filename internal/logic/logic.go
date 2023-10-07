package logic

import (
	"database/sql"
	"net/http"

	"github.com/solineun/ffcrm/internal/logic/logger"
	"github.com/solineun/ffcrm/internal/logic/storage"
	"github.com/solineun/ffcrm/internal/logic/handler"
	"github.com/solineun/ffcrm/pkg/logadapt"
	"github.com/solineun/ffcrm/pkg/models"
	"github.com/solineun/ffcrm/pkg/models/pg"
)

type Logic interface {
	logger.ErrLogger
	logger.InfoLogger
	storage.FFcrmDB
	handler.Handler
}

type LogicAdapter struct {
	log *logadapt.LoggerAdapter
	db  *pg.FFcrmDBadapt
}

// Home implements Handler.
func (la *LogicAdapter) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HOME PAGE\n"))
}

// GetOrderById implements FFcrmDB.
func (la *LogicAdapter) GetOrderById(productId int) (*models.Order, error) {
	order, err := la.db.GetOrderById(productId)
	return order, err
}

// InsertOrder implements FFcrmDB.
func (la *LogicAdapter) InsertOrder(productName string) (int, error) {
	id, err := la.db.InsertOrder(productName)
	return id, err
}

// LatestFiveOrders implements FFcrmDB.
func (la *LogicAdapter) LatestFiveOrders() ([]*models.Order, error) {
	orders, err := la.db.LatestFiveOrders()
	return orders, err
}

// Fatal implements ErrLogger.
func (la LogicAdapter) Fatal(err error) {
	la.log.Fatal(err)
}

// Printf implements InfoLogger.
func (la LogicAdapter) Printf(format string, v any) {
	la.log.Printf(format, v)
}

func NewLogicAdapter(db *sql.DB) *LogicAdapter {
	return &LogicAdapter{
		log: logadapt.NewLoggerAdapter(),
		db:  pg.NewFFcrmDB(db),
	}
}

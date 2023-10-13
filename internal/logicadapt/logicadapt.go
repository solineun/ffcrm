package logicadapt

import (
	"net/http"

	"github.com/solineun/ffcrm/internal/rbac"
	"github.com/solineun/ffcrm/internal/storage"
	la "github.com/solineun/ffcrm/pkg/loggeradapt"
)

type LogicAdapter struct {
	log la.LoggerAdapter
	db  storage.FFcrmDB
	rbac rbac.Rbac

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
		rbac: nil,
	}
}

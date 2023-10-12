package pg

import "github.com/solineun/ffcrm/pkg/models"

// GetOrderById implements storage.FFcrmDB.
func (db *FFcrmDBadapt) GetOrderById(productId int) (*models.Order, error) {
	
	panic("unimplemented")
}

// InsertOrder implements storage.FFcrmDB.
func (db *FFcrmDBadapt) InsertOrder(productName string) (int, error) {
	panic("unimplemented")
}

// LatestFiveOrders implements storage.FFcrmDB.
func (db *FFcrmDBadapt) LatestFiveOrders() ([]*models.Order, error) {
	
	panic("unimplemented")
}

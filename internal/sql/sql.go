package sql

import "github.com/solineun/ffcrm/pkg/models"

type FfcrmDb interface {
	Orders
}

type Orders interface {
	InsertOrder(productName string) (int, error)
	GetOrderById(productId int) (*models.Order, error)
	LatestFiveOrders() ([]*models.Order, error) 
}
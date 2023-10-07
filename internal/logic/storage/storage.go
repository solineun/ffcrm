package storage

import "github.com/solineun/ffcrm/pkg/models"

type FFcrmDB interface {
	orders
}

type orders interface {
	InsertOrder(productName string) (int, error)
	GetOrderById(productId int) (*models.Order, error)
	LatestFiveOrders() ([]*models.Order, error) 
}
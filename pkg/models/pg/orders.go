package pg

import (
	"database/sql"
	"solineun/ffcrm/pkg/models"
)

type OrderModel struct {
	DB *sql.DB
}

func (om *OrderModel) Insert(productName string) error {
	return nil
}

func (om *OrderModel) Get(id uint) (*models.Order, error) {
	return nil, nil
}

func (om *OrderModel) Latest() ([]*models.Order, error) {
	return nil, nil
}
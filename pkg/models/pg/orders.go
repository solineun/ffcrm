package pg

import (
	"database/sql"
	"errors"
	"solineun/ffcrm/pkg/models"
)

type OrderModel struct {
	DB *sql.DB
}

func (om *OrderModel) Insert(productName string) (int, error) {
	query := `INSERT INTO orders (product_name, created) 
	VALUES ($1, NOW()) RETURNING id`

	var id int
	err := om.DB.QueryRow(query, productName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (om *OrderModel) Get(id *int) (*models.Order, error) {
	query := `SELECT * FROM orders WHERE id = $1`	
	
	order := new(models.Order)
	err := om.DB.QueryRow(query, &id).Scan(&order.Id, &order.ProductName, &order.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return order, nil
}

func (om *OrderModel) Latest() ([]*models.Order, error) {
	return nil, nil
}

func (om *OrderModel) IsEmpty(o *models.Order) bool {
	return *o == models.Order{}
}
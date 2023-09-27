package pg

import (
	"database/sql"
	"solineun/ffcrm/pkg/models"
)

type OrderModel struct {
	DB *sql.DB
}

func (om *OrderModel) Insert(productName string) (uint, error) {
	query := `INSERT INTO orders (product_name, created) 
	VALUES ($1, NOW()) RETURNING id`

	var id int
	err := om.DB.QueryRow(query, productName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

func (om *OrderModel) Get(id uint) (*models.Order, error) {
	query := `SELECT * FROM orders WHERE id = $1`
	
	row := om.DB.QueryRow(query, id)	
	
	order := new(models.Order)
	err := row.Scan(order.Id, order.ProductName, order.Created)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (om *OrderModel) Latest() ([]*models.Order, error) {
	return nil, nil
}
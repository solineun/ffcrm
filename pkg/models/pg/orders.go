package pg

import (
	"database/sql"
	"errors"
	"github.com/solineun/ffcrm/pkg/models"

	"github.com/lib/pq"
)

type orderModel struct {}

func (om *orderModel) InsertOrder(productName string) (int, error) {
	query := `INSERT INTO orders (product_name, created) 
	VALUES ($1, NOW()) RETURNING id`

	var id int
	err := ffcrmDB.QueryRow(query, productName).Scan(&id)
	if err, ok := err.(*pq.Error); ok {
		if err.Code.Name() == "string_data_right_truncation" {
			return 0, models.ErrLongValue
		}
		return 0, err
	}

	return id, nil
}

func (om *orderModel) GetOrderById(productId int) (*models.Order, error) {
	query := `SELECT * FROM orders WHERE id = $1`	
	id := productId
	
	order := new(models.Order)
	err := ffcrmDB.QueryRow(query, &id).Scan(&order.Id, &order.ProductName, &order.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return order, nil
}

func (om *orderModel) LatestFiveOrders() ([]*models.Order, error) {
	query := `SELECT sub.* 
				FROM (SELECT * 
					FROM orders 
					ORDER BY created DESC 
					LIMIT 5
					) sub 
				ORDER BY created ASC`

	rows, err := ffcrmDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.Order
		
	for rows.Next() {
		o := new(models.Order)
		if err := rows.Scan(
			&o.Id, &o.ProductName, &o.Created); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return orders, models.ErrNoRecord
			}
			return orders, err
		}
		orders = append(orders, o)
	}
	if err = rows.Err(); err != nil {
		return orders, err
	}
	return orders, nil
}


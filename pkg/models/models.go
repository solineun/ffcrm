package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching entry found")
var ErrLongValue = errors.New("models: value too long")

type Order struct {
	Id string					`db:"id"`
	ProductId int			`db:"product_id"`
	CustomerId string		`db:"customer_id"`
	StockLocation string	`db:"stock_location"`
	CreatedAt time.Time		`db:"created_at"`
	Status Status			`db:"status"`
}

type Status string

const (
	PROCESSING Status = "ORDER_PROCESSING"
	ACCEPTED Status = "ORDER_ACCEPTED"
	//
	COMPLETED Status = "ORDER_COMPLETED"
	CANCELED Status = "ORDER_CANCELED"	
)

type Product struct {
	Id int				`db:"id"`
	Name string			`db:"name"`
	Description string	`db:"description"`
	Metrics Metrics		`db:"metrics"`
	CreatedAt time.Time	`db:"created_at"`
	UpdatedAt time.Time	`db:"updated_at"`
}

type Metrics struct {
	WidthMm int		`json: "width_mm"`
	LengthMm int	`json: "length_mm"`
	HeightMm int	`json: "height_mm"`
	WeightGr int	`json: "weight_gr"`
}

type User struct {
	Id string			`db:"id"`
	Login string		`db:"login"`
	Password string		`db:"password"`
	FirstName string	`db:"first_name"`
	SecondName string	`db:"second_name"`
	Email string		`db:"email"`
	PhoneNumber string	`db:"phone_number"`
	Role Role			`db:"role"`
}

type Role string

const (
	ADMIN Role = "ADMIN"
	CUSTOMER Role = "CUSTOMER"
)

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

type MarketsOrders struct {
	OrderId string
	MarketId int
}

type Market struct {
	Id int
	Name MarketName
}

type MarketName string

const (
	OZON MarketName = "OZON"
	WB MarketName = "WB"
)
package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching entry found")

type Order struct {
	Id uint
	ProductName string
	Created time.Time
}
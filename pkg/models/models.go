package models

import (
	"errors"
	"fmt"
	"time"
)

var ErrNoRecord = errors.New("models: no matching entry found")
var ErrLongValue = errors.New("models: value too long")

type Order struct {
	Id int
	ProductName string
	Created time.Time
}

func (o Order) Format() string {
	return fmt.Sprintf(
		"Oder id: %d\nProduct Name:" +  
		"%s\nCreated at: %s\n______________________\n", 
		o.Id, o.ProductName, o.Created.String())
}
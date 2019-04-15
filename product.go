package gorlick

import (
	"fmt"
)

// ProductUnit is an enum type representing the units for products.
type ProductUnit int

const (
	// THING is a general unit for the products.
	THING ProductUnit = 1 + iota
	// LITER is a general unit for a liquid products.
	LITER
	// GRAMM is a general unit for a solid products.
	GRAMM
)

var units = [...]string{
	"thing",
	"ml",
	"gr",
}

func (u ProductUnit) String() string {
	return units[u-1]
}

// Product is a model of all products.
type Product struct {
	ID          int
	Description string
	Name        string
	Unit        ProductUnit
}

func (p *Product) String() string {
	return fmt.Sprintf("Product [%d, %s, %s, %s]", p.ID, p.Name, p.Unit, p.Description)
}

// MakeProduct is a Product's factory.
func MakeProduct(ID int, Name string, Unit ProductUnit, Description string) *Product {
	p := &Product{
		ID:          ID,
		Name:        Name,
		Unit:        Unit,
		Description: Description,
	}
	return p
}

package gorlick

import (
	"fmt"
)

// ItemUnit is an enum type representing the units for items.
type ItemUnit int

const (
	// THING is a general unit for the items.
	THING ItemUnit = 1 + iota
	// LITER is a general unit for a liquid items.
	LITER
	// GRAMM is a general unit for a solid items.
	GRAMM
)

var units = [...]string{
	"thing",
	"ml",
	"gr",
}

func (u ItemUnit) String() string {
	return units[u-1]
}

// Item is a model of all products.
type Item struct {
	ID          int
	Description string
	Name        string
	Unit        ItemUnit
}

func (item *Item) String() string {
	return fmt.Sprintf("Item [%d, %s, %s, %s]", item.ID, item.Name, item.Unit, item.Description)
}

// MakeItem is a Items's factory.
func MakeItem(ID int, Name string, Unit ItemUnit, Description string) *Item {
	item := &Item{
		ID:          ID,
		Name:        Name,
		Unit:        Unit,
		Description: Description,
	}
	return item
}

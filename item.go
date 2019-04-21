package gorlick

import (
	"fmt"
)

// --- UNIT ---

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

// --- STATE ---

// ItemState is a state of an item.
type ItemState int

const (
	// FRIED is a state after frying action.
	FRIED ItemState = 1 + iota
)

// AddState adds a state to properties of the item.
func (item *Item) AddState(state ItemState) {
	item.States = append(item.States, state)
}

// --- ITEM ---

// Item is a model of all products.
type Item struct {
	ID          int
	Description string
	Name        string
	Unit        ItemUnit
	States      []ItemState
}

func (item *Item) String() string {
	return fmt.Sprintf("Item [%d, %s, %s, %s, %v]", item.ID, item.Name, item.Unit, item.Description, item.States)
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

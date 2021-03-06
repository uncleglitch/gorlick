package gorlick

import (
	"fmt"
)

// --- UNIT ---

// ItemUnit is an enum type representing the units for items.
type ItemUnit int

const (
	// THING is a general unit for the items.
	THING ItemUnit = iota
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
	return units[u]
}

// --- STATE ---

type void struct{}

var emptyMember void

// ItemState is a state of an item.
type ItemState int

const (
	// FRIED is a state after frying action.
	FRIED ItemState = iota
)

var states = [...]string{
	"fried",
}

func (s ItemState) String() string {
	return states[s]
}

// AddState adds a state to properties of the item.
func (item *Item) AddState(state ItemState) {
	item.States[state] = emptyMember
}

// GetStates returns all states of the item.
func (item *Item) GetStates() []ItemState {
	states := make([]ItemState, 0)
	for state := range item.States {
		states = append(states, state)
	}
	return states
}

// --- ITEM ---

// Item is a model of all products.
type Item struct {
	Description string
	Name        string
	Unit        ItemUnit
	States      map[ItemState]void
}

func (item *Item) String() string {
	return fmt.Sprintf("Item [%s, %s]", item.Name, item.Unit)
}

// MakeItem is a Items's factory.
func MakeItem(Name string, Unit ItemUnit) *Item {
	item := &Item{
		Name:        Name,
		Unit:        Unit,
		Description: "",
	}
	item.States = make(map[ItemState]void)
	return item
}

// --- ITEM'S BASE ---

// ItemsBase is a dictionary mapping item's name to an item.
var ItemsBase map[string]*Item

// AddItemToBase adds an item to the item's base.
func AddItemToBase(i *Item) {
	ItemsBase[i.Name] = i
}

// InitItemsBase initializes and loads the ItemsBase dictionary.
func InitItemsBase() {
	ItemsBase = make(map[string]*Item)
	AddItemToBase(MakeItem("Nil", THING))
}

// Copy copies the item.
func (item *Item) Copy() Item {
	copy := Item{
		Description: item.Description,
		Name:        item.Name,
		Unit:        item.Unit,
		States:      make(map[ItemState]void),
	}
	for state := range item.States {
		copy.AddState(state)
	}
	return copy
}

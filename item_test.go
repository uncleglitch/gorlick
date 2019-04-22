package gorlick_test

import (
	"testing"

	g "github.com/uncleglitch/gorlick"
)

func TestMakeItem(t *testing.T) {
	name := "TestName"
	unit := g.THING
	item := g.MakeItem(name, unit)
	if item.Name != name || item.Unit != unit {
		t.Error("MakeItem hasn't make right item.")
	}
}

func TestString(t *testing.T) {
	item := g.MakeItem("Egg", g.THING)
	if item.String() != "Item [Egg, thing]" {
		t.Error("String method for the item returns string with wrong format.")
	}
}

func TestAddState(t *testing.T) {
	item := g.MakeItem("Egg", g.THING)
	wantState := g.FRIED
	item.AddState(wantState)
	_, exists := item.States[wantState]
	if !exists {
		t.Error("New state hasn't added.")
	}
}

func TestInitItemBase(t *testing.T) {
	g.InitItemsBase()

	if len(g.ItemsBase) == 0 {
		t.Error("ItemsBase isn't initialized.")
	}
}

func TestAddItemToBase(t *testing.T) {
	g.InitItemsBase()

	countItemsAfterInitialize := len(g.ItemsBase)
	g.AddItemToBase(g.MakeItem("Test", g.THING))
	countItemsAfterAddingNewItem := len(g.ItemsBase)

	if countItemsAfterInitialize != countItemsAfterAddingNewItem-1 {
		t.Error("New item didn't add to the ItemsBase")
	}
}

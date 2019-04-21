package gorlick_test

import (
	"testing"

	g "github.com/uncleglitch/gorlick"
)

func TestMakeItem(t *testing.T) {
	name := "TestName"
	id := 1
	unit := g.THING
	description := "Test description"
	item := g.MakeItem(id, name, unit, description)
	if item.ID != id || item.Name != name || item.Unit != unit || item.Description != description {
		t.Error("MakeItem hasn't make right item.")
	}
}

func TestString(t *testing.T) {
	item := g.MakeItem(2, "Egg", g.THING, "A simple egg")
	if item.String() != "Item [2, Egg, thing, A simple egg]" {
		t.Error("String method for the item returns string with wrong format.")
	}
}

func TestAddState(t *testing.T) {
	item := g.MakeItem(2, "Egg", g.THING, "A simple egg")
	wantState := g.FRIED
	item.AddState(wantState)
	_, exists := item.States[wantState]
	if !exists {
		t.Error("New state hasn't added.")
	}
}

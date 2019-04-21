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
		t.Fail()
	}
}

func TestString(t *testing.T) {
	item := g.MakeItem(2, "Egg", g.THING, "A simple egg")
	if item.String() != "Item [2, Egg, thing, A simple egg]" {
		t.Fail()
	}
}

package gorlick_test

import (
	"testing"

	g "github.com/uncleglitch/gorlick"
)

func TestMakeProduct(t *testing.T) {
	name := "TestName"
	id := 1
	unit := g.THING
	description := "Test description"
	p := g.MakeProduct(id, name, unit, description)
	if p.ID != id || p.Name != name || p.Unit != unit || p.Description != description {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	p := g.MakeProduct(2, "Egg", g.THING, "A simple egg")
	if p.String() != "Product [2, Egg, thing, A simple egg]" {
		t.Fail()
	}
}

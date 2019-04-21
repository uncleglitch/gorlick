package gorlick_test

import (
	"reflect"
	"testing"

	g "github.com/uncleglitch/gorlick"
)

func TestMakeContainer(t *testing.T) {
	name := "container1"
	containerType := g.STANDARD
	p1 := g.MakeItem(0, "Item0", g.THING, "Description of Item0")
	p2 := g.MakeItem(1, "Item1", g.LITER, "Description of Item1")
	items := make([]g.Item, 2)
	items = append(items, *p1)
	items = append(items, *p2)

	c := g.MakeContainer(name, containerType, items)
	if c.Name != name || c.Type != containerType || !reflect.DeepEqual(c.Items, items) {
		t.Fail()
	}
}

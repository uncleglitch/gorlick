package gorlick_test

import (
	"reflect"
	"testing"

	g "github.com/uncleglitch/gorlick"
)

func TestMakeContainer(t *testing.T) {
	name := "container1"
	containerType := g.STANDARD
	p1 := g.MakeItem("Item0", g.THING)
	p2 := g.MakeItem("Item1", g.LITER)
	items := make([]g.ItemInContainer, 2)
	items = append(items, g.ItemInContainer{Item: p1, Count: 1})
	items = append(items, g.ItemInContainer{Item: p2, Count: 1})

	c := g.MakeContainer(name, containerType, items)
	if c.Name != name || c.Type != containerType || !reflect.DeepEqual(c.Items, items) {
		t.Error("MakeContainer hasn't make right container")
	}
}

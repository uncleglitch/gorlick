package gorlick_test

import (
	"reflect"
	"testing"
	"time"

	g "github.com/uncleglitch/gorlick"
)

func TestMakeAction(t *testing.T) {
	item1 := g.MakeItem(0, "Item0", g.THING, "Description of Item0")
	items := make([]g.Item, 2)
	items = append(items, *item1)
	c := g.MakeContainer("Container1", g.STANDARD, items)
	name := "Action1"
	description := "Description of action1"
	duration := time.Duration(10)
	a := g.MakeAction(name, c, description, duration)
	if a.Name != name || a.Container1 != c || a.Description != description || a.Duration != duration {
		t.Fail()
	}
}

func TestMove(t *testing.T) {
	item1 := g.MakeItem(0, "Item0", g.THING, "Description of Item0")
	items1 := make([]g.Item, 0)
	items1 = append(items1, *item1)
	c1 := g.MakeContainer("c1", g.STANDARD, items1)
	items2 := make([]g.Item, 0)
	c2 := g.MakeContainer("c2", g.STANDARD, items2)
	action := g.MakeAction("Move", c1, "Move action", 3*time.Second)
	action.Container2 = c2

	g.Move(action)

	if !reflect.DeepEqual(action.Container2.Items, items1) || len(action.Container1.Items) != 0 {
		t.Fail()
	}
}

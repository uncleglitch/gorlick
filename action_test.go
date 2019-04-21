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
	if a.Name != name || a.ContainerMain != c || a.Description != description || a.Duration != duration {
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
	action.ContainerHelp = c2

	g.Move(action)

	if !reflect.DeepEqual(action.ContainerHelp.Items, items1) || len(action.ContainerMain.Items) != 0 {
		t.Fail()
	}
}

func TestFry(t *testing.T) {
	item1 := g.MakeItem(0, "Item0", g.THING, "Description of Item0")
	items1 := make([]g.Item, 0)
	items1 = append(items1, *item1)
	c1 := g.MakeContainer("c1", g.STANDARD, items1)
	action := g.MakeAction("Fry", c1, "Fry action", 3*time.Minute)

	g.Fry(action)

	for _, item := range action.ContainerMain.Items {
		_, exists := item.States[g.FRIED]
		if !exists {
			t.Fail()
		}
	}
}

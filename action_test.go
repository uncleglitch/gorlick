package gorlick_test

import (
	"reflect"
	"testing"

	g "github.com/uncleglitch/gorlick"
)

func TestMakeAction(t *testing.T) {
	item1 := g.MakeItem("Item0", g.THING)
	items := make([]g.ItemInContainer, 0)
	items = append(items, g.ItemInContainer{Item: item1, Count: 1})
	c := g.MakeContainer("Container1", g.STANDARD, items)
	a := g.MakeAction(g.INIT, c)
	if a.ContainerMain != c {
		t.Error("MakeAction hasn't make right action.")
	}
}

func TestMove(t *testing.T) {
	item1 := g.MakeItem("Item0", g.THING)
	items1 := make([]g.ItemInContainer, 0)
	items1 = append(items1, g.ItemInContainer{Item: item1, Count: 1})
	c1 := g.MakeContainer("c1", g.STANDARD, items1)
	items2 := make([]g.ItemInContainer, 0)
	c2 := g.MakeContainer("c2", g.STANDARD, items2)
	action := g.MakeAction(g.MOVE, c1)
	action.ContainerHelp = c2

	g.Move(action)

	ok := false
	for _, sourceItem := range items1 {
		ok = false
		for _, targetItem := range action.ContainerHelp.Items {
			if reflect.DeepEqual(sourceItem, targetItem) {
				ok = true
				break
			}
		}
		if !ok {
			t.Error("The target container doesn't collect source's items.")
			break
		}
	}

	if len(action.ContainerMain.Items) != 0 {
		t.Error("The source container hasn't clear after moving.")
	}

}

func TestFry(t *testing.T) {
	item1 := g.MakeItem("Item0", g.THING)
	items1 := make([]g.ItemInContainer, 0)
	items1 = append(items1, g.ItemInContainer{Item: item1, Count: 1})
	c1 := g.MakeContainer("c1", g.STANDARD, items1)
	action := g.MakeAction(g.FRY, c1)

	g.Fry(action)

	for _, itemInContainer := range action.ContainerMain.Items {
		_, exists := itemInContainer.Item.States[g.FRIED]
		if !exists {
			t.Error("An item hasn't FRIED state")
		}
	}
}

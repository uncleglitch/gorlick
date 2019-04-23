package gorlick_test

import (
	"testing"

	g "github.com/uncleglitch/gorlick"
)

func TestMakeScenario(t *testing.T) {
	name := "scenario1"
	id := 0

	s := g.MakeScenario(id, name)

	if s.ID != id || s.Name != name {
		t.Error("The scenario was wrong initialized.")
	}

	anItem := g.MakeItem("Test", g.THING)
	s.AddItem(anItem)
	items := make([]g.ItemInContainer, 0)
	items = append(items, g.ItemInContainer{Item: anItem, Count: 1})
	c := g.MakeContainer("Test", g.STANDARD, items)
	s.AddContainer(c)
	s.AddAction(g.MakeAction(g.INIT, c))
}

func TestAddContainer(t *testing.T) {
	s := g.MakeScenario(0, "Test")
	containersCountBefore := len(s.Containers)
	s.AddContainer(g.MakeContainer("Test", g.STANDARD, make([]g.ItemInContainer, 0)))
	containersCountAfter := len(s.Containers)

	if containersCountAfter != containersCountBefore+1 {
		t.Error("AddContainer hasn't added a container into the scenario.")
	}
}

func TestAddContainers(t *testing.T) {
	s := g.MakeScenario(0, "Test")
	containersCountBefore := len(s.Containers)
	s.AddContainers([]*g.Container{
		g.MakeContainer("Test1", g.STANDARD, make([]g.ItemInContainer, 0)),
		g.MakeContainer("Test2", g.STANDARD, make([]g.ItemInContainer, 0)),
	})
	containersCountAfter := len(s.Containers)

	if containersCountAfter != containersCountBefore+2 {
		t.Error("AddContainers hasn't added the containers into the scenario.")
	}
}

func TestAddItem(t *testing.T) {
	s := g.MakeScenario(0, "Test")
	itemsCountBefore := len(s.Items)
	s.AddItem(g.MakeItem("Test", g.THING))
	itemsCountAfter := len(s.Items)

	if itemsCountAfter != itemsCountBefore+1 {
		t.Error("AddItem hasn't added an item into the scenario.")
	}
}

func TestAddItems(t *testing.T) {
	s := g.MakeScenario(0, "Test")
	itemsCountBefore := len(s.Items)
	s.AddItems([]*g.Item{
		g.MakeItem("Test1", g.THING),
		g.MakeItem("Test2", g.THING),
	})
	itemsCountAfter := len(s.Items)

	if itemsCountAfter != itemsCountBefore+2 {
		t.Error("AddItems hasn't added the items into the scenario.")
	}
}

func TestAddAction(t *testing.T) {
	s := g.MakeScenario(0, "Test")
	actionsCountBefore := len(s.Actions)
	s.AddAction(g.MakeAction(g.INIT, g.MakeContainer("Test", g.STANDARD, make([]g.ItemInContainer, 0))))
	actionsCountAfter := len(s.Actions)

	if actionsCountAfter != actionsCountBefore+1 {
		t.Error("AddAction hasn't added an action into the scenario.")
	}
}

package gorlick_test

import (
	"testing"

	g "github.com/uncleglitch/gorlick"
)

func TestRun(t *testing.T) {
	g.Init()

	s := g.MakeScenario(0, "TestScenario")

	i1 := g.MakeItem("Item1", g.THING)
	items1 := make([]g.ItemInContainer, 0)
	items1 = append(items1, g.ItemInContainer{Item: i1, Count: 1})
	c1 := g.MakeContainer("Container1", g.STANDARD, items1)

	i2 := g.MakeItem("Item2", g.THING)
	items2 := make([]g.ItemInContainer, 0)
	items2 = append(items2, g.ItemInContainer{Item: i2, Count: 1})
	c2 := g.MakeContainer("Container2", g.STANDARD, items2)

	s.AddContainer(c1)
	s.AddContainer(c2)
	s.AddItem(i1)
	s.AddItem(i2)

	a1 := g.MakeAction(g.MOVE, c1)
	a1.ContainerHelp = c2
	s.AddAction(a1)

	a2 := g.MakeAction(g.FRY, c2)
	s.AddAction(a2)

	_ = g.Run(s)
	//fmt.Println(m)
}

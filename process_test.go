package gorlick_test

import (
	"testing"
	"time"

	g "github.com/uncleglitch/gorlick"
)

func TestMakeProcessWithOneContainer(t *testing.T) {
	p1 := g.MakeProduct(0, "Product0", g.THING, "Description of Product0")
	products := make([]g.Product, 2)
	products = append(products, *p1)
	c := g.MakeContainer("Container1", g.STANDARD, products)
	name := "Process1"
	description := "Description of process1"
	duration := time.Duration(10)
	p := g.MakeProcessWithOneContainer(name, c, description, duration)
	if p.Name != name || p.ContainerIn1 != c || p.Description != description || p.Duration != duration {
		t.Fail()
	}
}

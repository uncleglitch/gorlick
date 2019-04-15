package gorlick_test

import (
	"reflect"
	"testing"

	g "github.com/uncleglitch/gorlick"
)

func TestMakeContainer(t *testing.T) {
	name := "container1"
	containerType := g.STANDARD
	p1 := g.MakeProduct(0, "Product0", g.THING, "Description of Product0")
	p2 := g.MakeProduct(1, "Product1", g.LITER, "Description of Product1")
	products := make([]g.Product, 2)
	products = append(products, *p1)
	products = append(products, *p2)

	c := g.MakeContainer(name, containerType, products)
	if c.Name != name || c.Type != containerType || !reflect.DeepEqual(c.Products, products) {
		t.Fail()
	}
}

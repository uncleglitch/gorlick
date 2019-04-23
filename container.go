package gorlick

import "fmt"

// ContainerType is an enum representing types of containers.
type ContainerType int

const (
	// STANDARD is the most general container.
	STANDARD ContainerType = iota
	// PLATE represents the all plates.
	PLATE
)

var containers = [...]string{
	"standard",
	"plate",
}

func (c ContainerType) String() string {
	return containers[c]
}

// ItemInContainer is an item with its count.
type ItemInContainer struct {
	Item  *Item
	Count float32
}

// Container is a general collection of items.
type Container struct {
	Name  string
	Items []ItemInContainer
	Type  ContainerType
}

func (c *Container) String() string {
	return fmt.Sprintf("Container %s [%v]", c.Type, c.Items)
}

// MakeContainer is the factory of Containers.
func MakeContainer(name string, containerType ContainerType, items []ItemInContainer) *Container {
	c := &Container{
		Name:  name,
		Type:  containerType,
		Items: items,
	}
	return c
}

// Add appends an item with its count to the container.
func (c *Container) Add(i *Item, count float32) {
	c.Items = append(c.Items, ItemInContainer{Item: i, Count: count})
}

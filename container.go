package gorlick

import "fmt"

// ContainerType is an enum representing types of containers.
type ContainerType int

const (
	// STANDARD is the most general container.
	STANDARD ContainerType = 1 + iota
	// PLATE represents the all plates.
	PLATE
)

var containers = [...]string{
	"standard",
	"plate",
}

func (c ContainerType) String() string {
	return containers[c-1]
}

// Container is a general collection of items.
type Container struct {
	Name  string
	Items []Item
	Type  ContainerType
}

func (c *Container) String() string {
	return fmt.Sprintf("Container %s [%v]", c.Type, c.Items)
}

// MakeContainer is the factory of Containers.
func MakeContainer(name string, containerType ContainerType, items []Item) *Container {
	c := &Container{
		Name:  name,
		Type:  containerType,
		Items: items,
	}
	return c
}

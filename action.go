package gorlick

import (
	"time"
)

// Action is a type representing the cook actions.
type Action struct {
	Name           string
	ContainerMain  *Container
	ContainerHelp  *Container
	ContainerTrash *Container
	Description    string
	Duration       time.Duration
}

// MakeAction makes a new action object.
func MakeAction(name string, container *Container) *Action {
	a := &Action{
		Name:          name,
		ContainerMain: container,
	}
	return a
}

// ActionType is a type of an action.
type ActionType int

const (
	// MOVE is a moving action.
	MOVE ActionType = 1 + iota
	// FRY is a frying action.
	FRY
)

var actions = [...]string{
	"move",
	"fry",
}

func (a ActionType) String() string {
	return actions[a]
}

// ActionFunction is a type representing an implementation of the actions.
type ActionFunction func(a *Action)

// MakeActionsBase makes dictionary maping ActionType to ActionFunction.
func MakeActionsBase() *map[ActionType]ActionFunction {

	actionsBase := map[ActionType]ActionFunction{
		MOVE: Move,
		FRY:  Fry,
	}
	return &actionsBase
}

// Move action moves items from containerMain to containerHelp.
func Move(a *Action) {
	for _, item := range a.ContainerMain.Items {
		a.ContainerHelp.Items = append(a.ContainerHelp.Items, item)
	}
	a.ContainerMain.Items = []ItemInContainer{}
}

// Fry action fries all items of main container.
func Fry(a *Action) {
	for _, itemInContainer := range a.ContainerMain.Items {
		itemInContainer.Item.AddState(FRIED)
	}
}

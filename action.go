package gorlick

import (
	"time"
)

// Action is a type representing the cook actions.
type Action struct {
	Name        string
	Container1  *Container
	Container2  *Container
	Container3  *Container
	Description string
	Duration    time.Duration
}

// MakeAction makes a new action object.
func MakeAction(name string, container *Container, description string, duration time.Duration) *Action {
	a := &Action{
		Name:        name,
		Container1:  container,
		Description: description,
		Duration:    duration,
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
		//FRY:  Fry,
	}
	return &actionsBase
}

// Move action moves items from container1 to container2.
func Move(a *Action) {
	for _, item := range a.Container1.Items {
		a.Container2.Items = append(a.Container2.Items, item)
	}
	a.Container1.Items = []Item{}
}

// func Fry(a *Action) {

// }

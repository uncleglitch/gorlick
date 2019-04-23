package gorlick

import (
	"time"
)

// Action is a type representing the cook actions.
type Action struct {
	Type           ActionType
	ContainerMain  *Container
	ContainerHelp  *Container
	ContainerTrash *Container
	Description    string
	Duration       time.Duration
}

// MakeAction makes a new action object.
func MakeAction(actionType ActionType, container *Container) *Action {
	a := &Action{
		Type:          actionType,
		ContainerMain: container,
	}
	a.ContainerHelp = new(Container)
	a.ContainerTrash = new(Container)
	return a
}

// ActionType is a type of an action.
type ActionType int

const (
	// INIT is a special action representing start of scenario.
	INIT ActionType = iota
	// END is a special action representing end of scenario.
	END
	// MOVE is a moving action.
	MOVE
	// FRY is a frying action.
	FRY
)

var actions = [...]string{
	"init",
	"end",
	"move",
	"fry",
}

func (a ActionType) String() string {
	return actions[a]
}

// ActionFunction is a type representing an implementation of the actions.
type ActionFunction func(a *Action)

// ActionsBase is a dictionary of the ActionFunctions.
var ActionsBase map[ActionType]ActionFunction

// InitActionsBase initializes dictionary maping ActionType to ActionFunction.
func InitActionsBase() {
	ActionsBase = map[ActionType]ActionFunction{
		INIT: Void,
		END:  Void,
		MOVE: Move,
		FRY:  Fry,
	}
}

// RunAction runs the implementation function of the action.
func (a *Action) RunAction() {
	ActionsBase[a.Type](a)
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

// Void is an empty function for the empty actions.
func Void(a *Action) {

}

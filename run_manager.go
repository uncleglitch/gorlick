package gorlick

import "fmt"

// ContainersDump is a slice of the all containers state after an action.
type ContainersDump struct {
	Containers map[string]Container
	Action     Action
}

// RunManager provides the running of the scenario.
type RunManager struct {
	Scenario *Scenario
	Dumps    []ContainersDump
}

// MakeContainerDump makes a slice of the all containers state in the scenario for the action.
func MakeContainerDump(s *Scenario, a *Action) ContainersDump {
	dump := ContainersDump{
		Containers: make(map[string]Container),
		Action:     *a,
	}
	dump.Containers = s.CopyContainers()
	return dump
}

// Run starts running of the scenario.
func Run(scenario *Scenario) *RunManager {
	m := &RunManager{
		Scenario: scenario,
		Dumps:    make([]ContainersDump, 0),
	}

	initAction := MakeAction(INIT, new(Container))
	m.Scenario.AddAction(initAction)
	m.Dumps = append(m.Dumps, MakeContainerDump(m.Scenario, initAction))

	for _, action := range m.Scenario.Actions {
		if action.Type != INIT {
			action.RunAction()
			m.Dumps = append(m.Dumps, MakeContainerDump(m.Scenario, action))
		}
	}

	m.Scenario.AddAction(MakeAction(END, new(Container)))

	return m
}

func (rm *RunManager) String() string {
	var s string
	for _, dump := range rm.Dumps {
		s = s + fmt.Sprintln(dump.Action.Type.String())
		for name, cont := range dump.Containers {
			s = s + fmt.Sprintf("\t[%s]:\n", name)
			for _, item := range cont.Items {
				s = s + fmt.Sprintf("\t\t%s (%.1f %s)\n", item.Item.Name, item.Count, item.Item.Unit)
			}
		}
	}
	return s
}

package gorlick

// Scenario is a sequence of the actions represents a recipe.
type Scenario struct {
	ID         int
	Name       string
	Actions    []*Action
	Containers map[string]*Container
	Items      map[string]*Item
}

// MakeScenario makes new scenario object.
func MakeScenario(id int, name string) *Scenario {
	var s = new(Scenario)
	s.ID = id
	s.Name = name
	s.Actions = make([]*Action, 0)
	s.Containers = make(map[string]*Container)
	s.Items = make(map[string]*Item)
	return s
}

// AddContainer adds a container to the scenario.
func (s *Scenario) AddContainer(c *Container) {
	s.Containers[c.Name] = c
}

// AddContainers adds the containers to the scenario.
func (s *Scenario) AddContainers(cs []*Container) {
	for _, c := range cs {
		s.Containers[c.Name] = c
	}
}

// AddItem adds an item to the scenario.
func (s *Scenario) AddItem(i *Item) {
	s.Items[i.Name] = i
}

// AddItems adds the items to the scenario.
func (s *Scenario) AddItems(is []*Item) {
	for _, i := range is {
		s.Items[i.Name] = i
	}
}

// AddAction adds an action to the scenario sequence.
func (s *Scenario) AddAction(a *Action) {
	s.Actions = append(s.Actions, a)
}

// CopyContainers creates a copy of all containers of the scenario.
func (s *Scenario) CopyContainers() map[string]Container {
	cs := make(map[string]Container)
	for name, container := range s.Containers {
		cs[name] = container.Copy()
	}
	return cs
}
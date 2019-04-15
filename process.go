package gorlick

import "time"

type Process struct {
	Name          string
	ContainerIn1  *Container
	ContainerIn2  *Container
	ContainerOut1 *Container
	ContainerOut2 *Container
	Description   string
	Duration      time.Duration
}

func MakeProcessWithOneContainer(name string, container *Container, description string, duration time.Duration) *Process {
	p := &Process{
		Name:         name,
		ContainerIn1: container,
		Description:  description,
		Duration:     duration,
	}
	return p
}

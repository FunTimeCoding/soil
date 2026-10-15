package context_change

import "github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"

type Change struct {
	Weight     *int32
	Payload    map[string]any
	Assignment *configuration_assignment.Assignment
}

package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Location(l *location.Location) *server.Location {
	return &server.Location{Identifier: l.Identifier, Name: l.Name}
}

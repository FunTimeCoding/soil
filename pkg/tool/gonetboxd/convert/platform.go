package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Platform(p *platform.Platform) *server.Platform {
	return &server.Platform{Identifier: p.Identifier, Name: p.Name}
}

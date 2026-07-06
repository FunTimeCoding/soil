package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Site(s *site.Site) *server.Site {
	return &server.Site{Identifier: s.Identifier, Name: s.Name}
}

package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/guest"
)

func (s *Server) guestLinks(f floor.Floor) map[guest.Guest]string {
	result := make(map[guest.Guest]string, len(f.Guests))

	for _, g := range f.Guests {
		result[g] = s.guestLink(g)
	}

	return result
}

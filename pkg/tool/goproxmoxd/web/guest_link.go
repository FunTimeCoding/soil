package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
)

func (s *Server) guestLink(g floor.Guest) string {
	i, okay := s.service.Instance(g.Hypervisor)

	if !okay {
		return ""
	}

	port := i.Port

	if port == 0 {
		port = constant.Port
	}

	return fmt.Sprintf(
		"https://%s:%d/#v1:0:=%s%%2F%d",
		i.Host,
		port,
		g.Kind,
		g.Identifier,
	)
}

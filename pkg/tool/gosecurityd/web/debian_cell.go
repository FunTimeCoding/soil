package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/flag_page"
	"maragu.dev/gomponents"
)

func (s *Server) debianCell(name string) gomponents.Node {
	return flag_page.DebianCell(s.debian.Status(name, ""))
}

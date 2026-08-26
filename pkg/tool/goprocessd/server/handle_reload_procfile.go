package server

import "fmt"

func (s *Server) handleReloadProcfile() string {
	if e := s.ReloadProcfile(); e != nil {
		return fmt.Sprintf("error: %s", e)
	}

	return "ok"
}

package server

import "fmt"

func (s *Server) handleReloadEnvironment() string {
	if e := s.ReloadEnvironment(); e != nil {
		return fmt.Sprintf("error: %s", e)
	}

	return "ok"
}

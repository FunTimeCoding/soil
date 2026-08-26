package server

import "fmt"

func (s *Server) handleRestart(arguments []string) string {
	if e := s.Restart(arguments); e != nil {
		return fmt.Sprintf("error: %s", e)
	}

	return "ok"
}

package server

import "fmt"

func (s *Server) handleRestartAll() string {
	count, e := s.RestartWave()

	if e != nil {
		return fmt.Sprintf("error: %s", e)
	}

	return fmt.Sprintf("ok: restarting %d processes in the background", count)
}

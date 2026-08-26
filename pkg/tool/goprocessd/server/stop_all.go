package server

func (s *Server) stopAll() error {
	var result error

	for _, p := range s.snapshotProcesses() {
		if e := p.Stop(); e != nil {
			result = e
		}
	}

	return result
}

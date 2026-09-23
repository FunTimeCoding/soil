package supervisor

func (s *Supervisor) stopAll() error {
	var result error

	for _, p := range s.snapshotProcesses() {
		if e := p.Stop(); e != nil {
			result = e
		}
	}

	return result
}

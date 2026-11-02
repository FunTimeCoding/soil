package supervisor

func (s *Supervisor) Start() {
	for _, p := range s.snapshotProcesses() {
		s.spawn(p)
	}
}

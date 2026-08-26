package server

func (s *Server) ReloadEnvironment() error {
	s.commandMutex.Lock()
	defer s.commandMutex.Unlock()

	return s.environment.Load(s.envrcPath)
}

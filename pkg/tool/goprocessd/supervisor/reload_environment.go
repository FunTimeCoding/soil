package supervisor

func (s *Supervisor) ReloadEnvironment() error {
	s.commandMutex.Lock()
	defer s.commandMutex.Unlock()

	return s.environment.Load(s.envrcPath)
}

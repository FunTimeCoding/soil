package scheduler

func (s *Scheduler) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.running {
		return
	}

	s.cron.Stop()
	s.running = false
}

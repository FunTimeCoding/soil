package turn

func (s *Scheduler) Waiting() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.queued()
}

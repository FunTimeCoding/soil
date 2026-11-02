package turn

func (s *Scheduler) release() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.order) == 0 {
		s.busy = false

		return
	}

	session := s.order[0]
	s.order = s.order[1:]
	queue := s.waiting[session]
	next := queue[0]

	if len(queue) > 1 {
		s.waiting[session] = queue[1:]
		s.order = append(s.order, session)
	} else {
		delete(s.waiting, session)
	}

	next.granted = true
	close(next.ready)
}

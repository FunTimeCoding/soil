package sweeper

import "time"

func (s *Sweeper) Start() {
	s.mutex.Lock()

	if s.running {
		s.mutex.Unlock()

		return
	}

	s.running = true
	s.done = make(chan struct{})
	s.mutex.Unlock()
	go s.run(time.NewTicker(s.interval))
}

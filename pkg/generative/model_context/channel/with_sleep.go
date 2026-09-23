package channel

import "time"

func (s *Server) WithSleep(f func(time.Duration)) *Server {
	s.sleep = f

	return s
}

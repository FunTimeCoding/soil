package base

func (s *Stack) Close() {
	if s.Client != nil {
		s.Client.Close()
	}

	s.Stop()
	s.Site.CloseClientConnections()
	s.Site.Close()
	s.pool.Close()
}

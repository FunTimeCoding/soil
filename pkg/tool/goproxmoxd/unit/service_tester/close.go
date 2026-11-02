package service_tester

func (s *Shells) Close() {
	for _, c := range s.Built {
		c.Close()
	}
}

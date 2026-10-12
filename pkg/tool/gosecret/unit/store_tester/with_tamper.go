package store_tester

func (s *Tester) WithTamper(key string) *Tester {
	s.tamper = key

	return s
}

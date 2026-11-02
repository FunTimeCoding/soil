package double

func (s *Set) Close() {
	s.MemoryPeer.Close()
	s.Query.Close()
}

package peer_tester

func (s *Peers) Close() {
	s.MemoryPeer.Close()
	s.Query.Close()
}

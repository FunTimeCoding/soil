package cruise

func (s *Service) releaseHeld(identifier string) (bool, error) {
	wakes, e := s.wakes(identifier)

	if e != nil || !wakes {
		return false, e
	}

	room, f := s.store.ImmediateRoom(identifier, "", s.clock())

	if f != nil || !room {
		return false, f
	}

	count, g := s.store.ReleaseHeldCruise(identifier)

	if g != nil || count == 0 {
		return false, g
	}

	return true, s.store.MarkWoken(identifier)
}

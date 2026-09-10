package store

import "math/rand"

func (s *Store) NextName() (string, error) {
	available, e := s.AvailableNames()

	if e != nil {
		return "", e
	}

	if len(available) == 0 {
		return "", nil
	}

	return available[rand.Intn(len(available))], nil
}

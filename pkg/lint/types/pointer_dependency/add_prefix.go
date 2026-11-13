package pointer_dependency

func (s *Set) AddPrefix(p string) {
	s.prefixes[p] = true
}

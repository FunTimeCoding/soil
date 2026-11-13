package pointer_dependency

func (s *Set) AddPath(p string) {
	s.paths[p] = true
}

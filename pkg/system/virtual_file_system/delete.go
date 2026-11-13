package virtual_file_system

func (s *System) Delete(path string) {
	if _, present := s.files[path]; present {
		s.sorted = nil
	}

	delete(s.files, path)
	s.deleted[path] = true
}

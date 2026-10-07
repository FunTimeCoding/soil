package sink

func (s *Sink) Paths() []string {
	var result []string

	for _, o := range s.operations {
		result = append(result, o.Path)
	}

	return result
}

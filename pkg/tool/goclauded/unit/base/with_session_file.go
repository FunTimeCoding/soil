package base

func WithSessionFile(
	identifier string,
	slug string,
) Option {
	return func(s *setup) {
		s.sessionFiles = append(s.sessionFiles, [2]string{identifier, slug})
	}
}

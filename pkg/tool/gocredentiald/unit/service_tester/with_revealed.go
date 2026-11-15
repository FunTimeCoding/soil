package service_tester

func WithRevealed(fields ...string) Option {
	return func(s *Tester) {
		s.revealed = fields
	}
}

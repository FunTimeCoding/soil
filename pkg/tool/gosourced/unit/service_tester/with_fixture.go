package service_tester

func WithFixture(
	name string,
	modLines ...string,
) Option {
	return func(s *setup) {
		s.fixture = name
		s.modLines = modLines
	}
}

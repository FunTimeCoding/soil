package service

import "context"

func (s *Service) Generate(
	x context.Context,
	names []string,
) error {
	return s.parserGuard.Attempt(
		x,
		func(y context.Context) error {
			_, e := s.parser.PostGenerate(y, s.EliteFiles(names), nil)

			return e
		},
	)
}

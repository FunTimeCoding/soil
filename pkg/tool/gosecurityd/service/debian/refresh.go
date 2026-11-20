package debian

import "context"

func (s *Service) Refresh(x context.Context) []error {
	var result []error

	for _, source := range s.sources {
		if e := source.Refresh(x); e != nil {
			result = append(result, e)
		}
	}

	return result
}

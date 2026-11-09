package service

import "time"

func (s *Service) CreateImpression(
	content string,
	source string,
	capturedAt string,
) (int64, error) {
	at := s.clock().UTC()

	if capturedAt != "" {
		parsed, e := time.Parse(time.RFC3339, capturedAt)

		if e != nil {
			return 0, e
		}

		at = parsed.UTC()
	}

	return s.store.CreateImpression(content, source, at)
}

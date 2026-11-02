package event

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) Summarize(
	sessionIdentifier string,
	name string,
	body string,
) error {
	if e := s.store.UpsertSummary(sessionIdentifier, name, body); e != nil {
		return e
	}

	if e := s.store.UpsertEvent(
		sessionIdentifier,
		constant.Summarize,
		name,
		map[string]string{constant.Body: body},
	); e != nil {
		return e
	}

	s.notify()
	slug, metadata, f := s.document.SessionMetadata(sessionIdentifier)

	if f != nil {
		return f
	}

	if slug == "" {
		return nil
	}

	s.report(s.document.PushSummary(slug, body, metadata))

	return nil
}

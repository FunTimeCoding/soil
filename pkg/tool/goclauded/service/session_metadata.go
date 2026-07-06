package service

import "github.com/funtimecoding/soil/pkg/time"

func (s *Service) sessionMetadata(
	sessionIdentifier string,
) (string, map[string]string, error) {
	session, e := s.store.GetSession(sessionIdentifier)

	if e != nil {
		return "", nil, e
	}

	if session == nil {
		return "", nil, nil
	}

	name := session.Slug

	if session.Alias != nil && *session.Alias != "" {
		name = *session.Alias
	}

	if name == "" {
		return "", nil, nil
	}

	result := map[string]string{
		"session_name": name,
		"session_id":   sessionIdentifier,
	}

	if !session.StartedAt.IsZero() {
		result["date"] = session.StartedAt.Format(time.DateYear)
	}

	return name, result, nil
}

package selection

import "context"

func (s *Selection) Active(x context.Context) (string, error) {
	session, e := identifier(x)

	if e != nil {
		return "", e
	}

	v, okay := s.sessions.Load(session)

	if !okay {
		return "", nil
	}

	return v.(string), nil
}

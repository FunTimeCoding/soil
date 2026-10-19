package selection

import "context"

func (s *Selection) Select(
	x context.Context,
	name string,
) error {
	session, e := identifier(x)

	if e != nil {
		return e
	}

	s.sessions.Store(session, name)

	return nil
}

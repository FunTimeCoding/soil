package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) Moment(
	sessionIdentifier string,
	name string,
	line string,
) error {
	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Moment,
		name,
		map[string]string{"line": line},
	); e != nil {
		return e
	}

	s.ForwardImpression(line, name)
	s.notify()

	return nil
}

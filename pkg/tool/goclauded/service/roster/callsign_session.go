package roster

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Service) CallsignSession(c string) (*session.Session, error) {
	return s.store.CallsignSession(c)
}

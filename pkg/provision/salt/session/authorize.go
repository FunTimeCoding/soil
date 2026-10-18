package session

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"net/http"
)

func (s *Session) Authorize(r *http.Request) error {
	if s.Token() == "" {
		if e := s.Login(r.Context()); e != nil {
			return e
		}
	}

	r.Header.Set(constant.SaltTokenHeader, s.Token())

	return nil
}

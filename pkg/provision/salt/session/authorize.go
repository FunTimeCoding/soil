package session

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"net/http"
)

func (s *Session) Authorize(r *http.Request) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	r.Header.Set(constant.SaltTokenHeader, s.token)

	return nil
}

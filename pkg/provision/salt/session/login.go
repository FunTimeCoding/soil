package session

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (s *Session) Login(x context.Context) error {
	var r response.Login

	if e := s.login.Notation(
		request.New(http.MethodPost, constant.SaltLoginPath).
			WithContext(x).
			WithNotation(s.request),
		&r,
	); e != nil {
		return e
	}

	if len(r.Return) == 0 || r.Return[0].Token == "" {
		return unexpected.Format("salt login answered without a token")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.token = r.Return[0].Token

	return nil
}

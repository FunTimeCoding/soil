package fake_assistant

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/url"
)

func (s *Server) Host() string {
	u, e := url.Parse(s.server.URL)
	errors.PanicOnError(e)

	return u.Host
}

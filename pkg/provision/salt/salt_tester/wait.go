package salt_tester

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"strings"
)

func (s *Server) wait(r *http.Request) {
	s.mutex.Lock()
	held, hold, holding := s.held, s.hold, s.holding
	s.mutex.Unlock()

	if hold == nil || strings.TrimPrefix(
		r.URL.Path,
		webConstant.RootPath,
	) != held {
		return
	}

	select {
	case holding <- struct{}{}:
	default:
	}

	<-hold
}

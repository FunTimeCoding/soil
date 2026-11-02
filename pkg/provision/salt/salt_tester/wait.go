package salt_tester

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"strings"
)

func (s *Server) wait(r *http.Request) {
	if strings.TrimPrefix(r.URL.Path, webConstant.RootPath) != constant.SaltKeysPath {
		return
	}

	s.mutex.Lock()
	hold, holding := s.hold, s.holding
	s.mutex.Unlock()

	if hold == nil {
		return
	}

	select {
	case holding <- struct{}{}:
	default:
	}

	<-hold
}

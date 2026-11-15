package proxmox_server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/proxmox/unit/proxmox_server/answer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
	"strings"
)

func (s *Server) serve(
	w http.ResponseWriter,
	r *http.Request,
) {
	path := strings.TrimPrefix(r.URL.Path, constant.Base)
	s.mutex.Lock()
	a, okay := s.answers[join.Space(r.Method, path)]

	if !okay {
		a, okay = s.answers[path]
	}

	s.mutex.Unlock()

	if !okay {
		a = answer.New("501 no such uri", "")
	}

	c, i, e := http.NewResponseController(w).Hijack()

	if e != nil {
		s.t.Errorf("hijack: %v", e)

		return
	}

	defer errors.LogClose(c)
	_, e = fmt.Fprintf(
		i,
		"HTTP/1.1 %s\r\nContent-Type: application/json\r\nContent-Length: %d\r\nConnection: close\r\n\r\n%s",
		a.Line,
		len(a.Body),
		a.Body,
	)
	errors.LogOnError(e)
	errors.LogOnError(i.Flush())
}

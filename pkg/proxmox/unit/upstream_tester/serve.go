package upstream_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
	"strings"
)

func (u *Upstream) serve(
	w http.ResponseWriter,
	q *http.Request,
) {
	path := strings.TrimPrefix(q.URL.Path, constant.Base)
	u.mutex.Lock()
	a, okay := u.answers[join.Space(q.Method, path)]

	if !okay {
		a, okay = u.answers[path]
	}

	u.mutex.Unlock()

	if !okay {
		a = &Answer{Line: "501 no such uri"}
	}

	c, b, e := http.NewResponseController(w).Hijack()

	if e != nil {
		u.t.Errorf("hijack: %v", e)

		return
	}

	defer errors.LogClose(c)
	_, e = fmt.Fprintf(
		b,
		"HTTP/1.1 %s\r\nContent-Type: application/json\r\nContent-Length: %d\r\nConnection: close\r\n\r\n%s",
		a.Line,
		len(a.Body),
		a.Body,
	)
	errors.LogOnError(e)
	errors.LogOnError(b.Flush())
}

package stream

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"io"
	"net/http"
)

func (s *Stream) Open(x context.Context) error {
	q, e := http.NewRequestWithContext(x, http.MethodGet, s.locator, nil)

	if e != nil {
		return e
	}

	web.Bearer(q, s.token)

	if s.last != "" {
		q.Header.Set(webConstant.LastEvent, s.last)
	}

	r, e := s.pool.Client().Do(q)

	if e != nil {
		return e
	}

	if r.StatusCode != http.StatusOK {
		defer errors.LogClose(r.Body)
		b, f := io.ReadAll(r.Body)

		if f != nil {
			return f
		}

		return refusal.New(constant.StreamOperation, r, b)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.body = r.Body

	return nil
}

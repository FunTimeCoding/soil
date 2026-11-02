package refusal_transport

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"net/http"
)

func (t *Transport) RoundTrip(q *http.Request) (*http.Response, error) {
	r, e := t.next.RoundTrip(q)

	if e != nil {
		if f := connection.Classify(e); f != nil {
			if f.Host == "" {
				return nil, connection.New(f.Kind, t.host, f.Path, f.Reason)
			}

			return nil, f
		}

		return nil, e
	}

	if r.StatusCode < http.StatusBadRequest ||
		r.StatusCode == http.StatusUnauthorized {
		return r, nil
	}

	defer errors.LogClose(r.Body)

	return nil, t.refuse(q, r)
}

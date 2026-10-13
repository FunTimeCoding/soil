package bearer_transport

import (
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (t *Transport) RoundTrip(e *http.Request) (*http.Response, error) {
	q := e.Clone(e.Context())
	q.Header.Set(
		constant.Authorization,
		key_value.Space(constant.Bearer, t.token),
	)

	return t.r.RoundTrip(q)
}

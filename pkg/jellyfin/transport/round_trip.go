package transport

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	q := r.Clone(r.Context())
	q.Header.Set(
		constant.Authorization,
		fmt.Sprintf("MediaBrowser Token=%q", t.token),
	)

	return t.base.RoundTrip(q)
}

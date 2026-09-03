package bearer

import (
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (Transport) RoundTrip(q *http.Request) (*http.Response, error) {
	q.Header.Set(
		web.Authorization,
		key_value.Space(web.Bearer, generative.ModelContextTestToken),
	)

	return http.DefaultTransport.RoundTrip(q)
}

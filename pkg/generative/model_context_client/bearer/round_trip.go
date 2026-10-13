package bearer

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (t Transport) RoundTrip(q *http.Request) (*http.Response, error) {
	q.Header.Set(
		webConstant.Authorization,
		key_value.Space(webConstant.Bearer, constant.ModelContextTestToken),
	)

	return t.Base.RoundTrip(q)
}

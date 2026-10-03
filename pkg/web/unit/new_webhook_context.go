package unit

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/request_context"
	"net/http"
	"net/http/httptest"
	"strings"
)

func newWebhookContext() *request_context.Context {
	q := httptest.NewRequest(
		http.MethodPost,
		"/hook",
		strings.NewReader("payload"),
	)
	q.Header.Set(constant.Authorization, "Bearer secret")

	return request_context.New(httptest.NewRecorder(), q)
}

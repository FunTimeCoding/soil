package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerSideEventOpensWithoutWrites(t *testing.T) {
	server := httptest.NewServer(
		layout.HandleServerSideEvent(
			notifier.New(),
			func(
				http.ResponseWriter,
				http.Flusher,
			) {
			},
		),
	)
	defer server.Close()
	response, e := http.Get(server.URL)
	errors.PanicOnError(e)
	defer errors.PanicClose(response.Body)
	assert.Integer(t, http.StatusOK, response.StatusCode)
	assert.String(
		t,
		"text/event-stream",
		response.Header.Get(constant.ContentType),
	)
}

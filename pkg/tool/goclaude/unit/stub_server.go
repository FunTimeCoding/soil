package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func stubServer(
	t *testing.T,
	body string,
	statuses ...int,
) *client.ClientWithResponses {
	t.Helper()
	call := 0
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				status := http.StatusOK

				if call < len(statuses) {
					status = statuses[call]
				}

				call++

				if status != http.StatusOK {
					w.WriteHeader(status)

					return
				}

				w.Header().Set("Content-Type", "application/json")
				_, e := w.Write([]byte(body))
				assert.FatalOnError(t, e)
			},
		),
	)
	t.Cleanup(s.Close)
	c, e := client.NewClientWithResponses(s.URL)
	assert.FatalOnError(t, e)

	return c
}

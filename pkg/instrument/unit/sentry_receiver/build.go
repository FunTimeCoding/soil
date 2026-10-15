package sentry_receiver

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/getsentry/sentry-go"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func build(
	t *testing.T,
	delay time.Duration,
) *atomic.Int32 {
	t.Helper()
	result := &atomic.Int32{}
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				_ http.ResponseWriter,
				q *http.Request,
			) {
				if strings.HasSuffix(q.URL.Path, "/envelope/") {
					time.Sleep(delay)
					result.Add(1)
				}
			},
		),
	)
	t.Cleanup(s.Close)
	t.Setenv(
		constant.LocatorEnvironment,
		fmt.Sprintf("http://key@%s/1", strings.TrimPrefix(s.URL, "http://")),
	)
	t.Cleanup(func() { sentry.CurrentHub().BindClient(nil) })

	return result
}

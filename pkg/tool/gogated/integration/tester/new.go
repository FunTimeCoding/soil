package tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/integration/base"
	"net/http"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)

	return &Tester{
		server: s,
		client: &http.Client{
			CheckRedirect: func(
				_ *http.Request,
				_ []*http.Request,
			) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

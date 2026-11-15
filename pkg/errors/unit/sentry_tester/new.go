package sentry_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()
	result := basic.New(
		locator.NewOrigin(server).Base(constant.Base).Trail(),
		"alfa-token",
	)
	t.Cleanup(result.Close)

	return result
}

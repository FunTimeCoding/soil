package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newSentryClient(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()

	return basic.New(
		upstream_tester.Locator(t, server).Base(constant.Base).Trail(),
		"alfa-token",
	)
}

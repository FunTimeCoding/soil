package unit

import (
	"github.com/funtimecoding/soil/pkg/opnsense/basic"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newClient(
	t *testing.T,
	server string,
	untrusted bool,
) *basic.Client {
	t.Helper()

	return basic.New(
		upstream_tester.Locator(t, server).Base(constant.Base),
		"alfa-key",
		"bravo-secret",
		untrusted,
	)
}

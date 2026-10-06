package unit

import (
	"github.com/funtimecoding/soil/pkg/habitica/basic"
	"github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newClient(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()

	return basic.New(
		upstream_tester.Locator(t, server).Base(constant.Base),
		"alfa-user",
		"bravo-token",
	)
}

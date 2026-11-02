package unit

import (
	"github.com/funtimecoding/soil/pkg/linkace/basic"
	"github.com/funtimecoding/soil/pkg/linkace/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newClient(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()
	result := basic.New(
		upstream_tester.Locator(t, server).Base(constant.BasePath),
		"alfa-token",
	)
	t.Cleanup(result.Close)

	return result
}

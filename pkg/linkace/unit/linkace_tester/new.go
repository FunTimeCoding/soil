package linkace_tester

import (
	"github.com/funtimecoding/soil/pkg/linkace/basic"
	"github.com/funtimecoding/soil/pkg/linkace/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()
	result := basic.New(
		locator.NewOrigin(server).Base(constant.BasePath),
		"alfa-token",
	)
	t.Cleanup(result.Close)

	return result
}

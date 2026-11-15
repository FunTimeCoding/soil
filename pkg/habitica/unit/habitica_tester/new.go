package habitica_tester

import (
	"github.com/funtimecoding/soil/pkg/habitica/basic"
	"github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()
	result := basic.New(
		locator.NewOrigin(server).Base(constant.Base),
		"alfa-user",
		"bravo-token",
	)
	t.Cleanup(result.Close)

	return result
}

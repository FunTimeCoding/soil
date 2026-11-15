package opnsense_tester

import (
	"github.com/funtimecoding/soil/pkg/opnsense/basic"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func Basic(
	t *testing.T,
	server string,
	untrusted bool,
) *basic.Client {
	t.Helper()
	result := basic.New(
		locator.NewOrigin(server).Base(constant.Base),
		"alfa-key",
		"bravo-secret",
		untrusted,
	)
	t.Cleanup(result.Close)

	return result
}

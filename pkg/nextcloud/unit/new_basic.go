package unit

import (
	"github.com/funtimecoding/soil/pkg/nextcloud/basic"
	"github.com/funtimecoding/soil/pkg/nextcloud/helper"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newBasic(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()
	host, p := upstream_tester.HostPort(t, server)

	return basic.New(
		helper.FileRootLocator(host, "alfa").Insecure().Port(p),
		"alfa",
		"bravo-password",
	)
}

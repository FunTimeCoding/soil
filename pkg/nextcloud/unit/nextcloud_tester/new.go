package nextcloud_tester

import (
	"github.com/funtimecoding/soil/pkg/nextcloud/basic"
	"github.com/funtimecoding/soil/pkg/nextcloud/helper"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func New(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()
	host, p := upstream_server.HostPort(t, server)
	result := basic.New(
		helper.FileRootLocator(host, "alfa").Insecure().Port(p),
		"alfa",
		"bravo-password",
	)
	t.Cleanup(result.Close)

	return result
}

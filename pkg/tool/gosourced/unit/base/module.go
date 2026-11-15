package base

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func Module(
	t *testing.T,
	directory string,
	token string,
) *client.Client {
	t.Helper()
	d, _ := daemon_tester.New(t, service_tester.Inventory(directory))

	return connect(t, d, token)
}

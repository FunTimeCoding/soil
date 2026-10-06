package unit

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/basic"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newSaltClient(
	t *testing.T,
	server string,
	password string,
) *basic.Client {
	t.Helper()

	return basic.New(
		upstream_tester.Locator(t, server),
		"alfa",
		password,
		"pam",
	)
}

package salt_tester

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/basic"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	server string,
	password string,
) *basic.Client {
	t.Helper()
	result := basic.New(locator.NewOrigin(server), "alfa", password, "pam")
	t.Cleanup(result.Close)

	return result
}

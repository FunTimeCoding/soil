package netbox_tester

import (
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	address string,
) *netbox.Client {
	t.Helper()
	result := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(address),
	)
	t.Cleanup(result.Close)

	return result
}

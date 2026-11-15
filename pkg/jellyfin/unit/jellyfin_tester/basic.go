package jellyfin_tester

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/basic"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func Basic(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()
	result := basic.New(locator.NewOrigin(server), "alfa-token")
	t.Cleanup(result.Close)

	return result
}

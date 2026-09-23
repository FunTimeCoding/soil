//go:build browser

package lifetime_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"testing"
)

func ProtocolAbsent(
	t *testing.T,
	c *chromium.Client,
) (result string) {
	t.Helper()

	defer func() {
		result = fmt.Sprint(recover())
	}()
	protocol.New(c, constant.FixtureAbsentTab)

	return ""
}

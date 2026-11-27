//go:build browser

package lifetime_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/browser/chromium"
	"github.com/funtimecoding/soil/pkg/browser/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/browser/constant"
)

func (o *Tester) ProtocolAbsent(c *chromium.Client) (result string) {
	o.t.Helper()

	defer func() {
		result = fmt.Sprint(recover())
	}()
	protocol.New(c, constant.ChromiumFixtureAbsentTab)

	return ""
}

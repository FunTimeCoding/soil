//go:build browser

package lifetime_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
)

func (o *Tester) ProtocolAbsent(c *chromium.Client) (result string) {
	o.t.Helper()

	defer func() {
		result = fmt.Sprint(recover())
	}()
	protocol.New(c, constant.FixtureAbsentTab)

	return ""
}

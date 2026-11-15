package cross_service_tester

import (
	chatConstant "github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
)

func (o *Tester) Eventually(condition func() bool) {
	o.t.Helper()
	wait_tester.Until(o.t, condition, chatConstant.MattermostSocketWait)
}

package cross_service_tester

import (
	chatConstant "github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"
	"time"
)

func (o *Tester) AwaitNotification() []notification.Notification {
	o.t.Helper()
	var result []notification.Notification
	deadline := time.Now().Add(chatConstant.MattermostSocketWait)

	for time.Now().Before(deadline) {
		result = o.Goclauded.Store.PendingNotifications("kilo-session")

		if len(result) > 0 {
			return result
		}

		time.Sleep(5 * time.Millisecond)
	}

	return result
}

package worker

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/integration/cross_service_tester"
	"testing"
	"time"
)

func waitForNotification(
	t *testing.T,
	r *cross_service_tester.Tester,
) []notification.Notification {
	t.Helper()
	var result []notification.Notification
	deadline := time.Now().Add(constant.MattermostSocketWait)

	for time.Now().Before(deadline) {
		result = r.Goclauded.Store.PendingNotifications("kilo-session")

		if len(result) > 0 {
			return result
		}

		time.Sleep(5 * time.Millisecond)
	}

	return result
}

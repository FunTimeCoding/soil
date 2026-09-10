package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"
)

func (o *Tester) PendingNotifications(
	sessionIdentifier string,
) []notification.Notification {
	result, e := o.Store.PendingNotifications(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}

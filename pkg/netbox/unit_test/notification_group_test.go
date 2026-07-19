package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/notification_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestNotificationGroup(t *testing.T) {
	assert.NotNil(t, notification_group.New(&netbox.NotificationGroup{}))
}

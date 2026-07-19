package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/power_feed"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestFeed(t *testing.T) {
	assert.NotNil(t, power_feed.New(&netbox.PowerFeed{}))
}

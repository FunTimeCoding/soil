package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_link"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestWirelessLink(t *testing.T) {
	assert.NotNil(t, wireless_link.New(&netbox.WirelessLink{}))
}

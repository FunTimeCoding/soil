package wireless_link

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestLink(t *testing.T) {
	assert.NotNil(t, New(&netbox.WirelessLink{}))
}

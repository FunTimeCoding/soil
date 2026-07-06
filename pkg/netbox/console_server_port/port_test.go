package console_server_port

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestPort(t *testing.T) {
	assert.NotNil(t, New(&netbox.ConsoleServerPort{}))
}

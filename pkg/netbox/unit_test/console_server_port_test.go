package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/console_server_port"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestConsoleServerPort(t *testing.T) {
	assert.NotNil(t, console_server_port.New(&netbox.ConsoleServerPort{}))
}

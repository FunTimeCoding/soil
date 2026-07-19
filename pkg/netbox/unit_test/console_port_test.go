package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/console_port"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestConsolePort(t *testing.T) {
	assert.NotNil(t, console_port.New(&netbox.ConsolePort{}))
}

//go:build local

package client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/opnsense"
	"testing"
)

func TestInterfacesAndLogRead(t *testing.T) {
	c := opnsense.NewEnvironment()
	interfaces, e := c.Interfaces()
	assert.FatalOnError(t, e)
	assert.NotEmpty(t, interfaces)
	entries, f := c.Log(1)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, entries)
}

//go:build local

package client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/technitium"
	"testing"
)

func TestZonesRead(t *testing.T) {
	zones, e := technitium.NewEnvironment().Zones()
	assert.FatalOnError(t, e)
	assert.NotEmpty(t, zones)
}

func TestMissingZoneIsRefused(t *testing.T) {
	_, e := technitium.NewEnvironment().Records("charlie.invalid", true)
	assert.True(t, unexpected.Is(e))
}

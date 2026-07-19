package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestPlatform(t *testing.T) {
	assert.NotNil(t, platform.New(&netbox.Platform{}))
}

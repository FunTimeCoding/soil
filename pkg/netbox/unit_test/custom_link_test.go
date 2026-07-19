package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/custom_link"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestCustomLink(t *testing.T) {
	assert.NotNil(t, custom_link.New(&netbox.CustomLink{}))
}

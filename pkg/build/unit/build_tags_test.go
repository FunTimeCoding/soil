package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/build"
	"testing"
)

func TestTagsCarryTimeZonesWithoutRequest(t *testing.T) {
	assert.String(t, "timetzdata", build.Tags(""))
}

func TestTagsKeepTimeZonesBesideRequestedTags(t *testing.T) {
	assert.String(t, "timetzdata,netgo", build.Tags("netgo"))
}

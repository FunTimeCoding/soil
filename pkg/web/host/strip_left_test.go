package host

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestStripLeft(t *testing.T) {
	assert.String(t, "", StripLeft(""))
	assert.String(t, "org", StripLeft("org"))
	assert.String(t, "org", StripLeft(constant.Example))
	assert.String(t, "example.org", StripLeft("test.example.org"))
	assert.String(
		t,
		"test.example.org",
		StripLeft("test2.test.example.org"),
	)
}

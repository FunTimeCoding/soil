package host

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestStripDomain(t *testing.T) {
	assert.String(t, "", StripDomain(""))
	assert.String(t, "org", StripDomain("org"))
	assert.String(t, "example", StripDomain(constant.Example))
	assert.String(
		t,
		"test",
		StripDomain("test.example.org"),
	)
	assert.String(
		t,
		"test2",
		StripDomain("test2.test.example.org"),
	)
}

package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/host"
	"testing"
)

func TestStripDomain(t *testing.T) {
	assert.String(t, "", host.StripDomain(""))
	assert.String(t, "org", host.StripDomain("org"))
	assert.String(t, "example", host.StripDomain(constant.Example))
	assert.String(
		t,
		"test",
		host.StripDomain("test.example.org"),
	)
	assert.String(
		t,
		"test2",
		host.StripDomain("test2.test.example.org"),
	)
}

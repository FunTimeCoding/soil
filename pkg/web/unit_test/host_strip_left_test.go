package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/host"
	"testing"
)

func TestStripLeft(t *testing.T) {
	assert.String(t, "", host.StripLeft(""))
	assert.String(t, "org", host.StripLeft("org"))
	assert.String(t, "org", host.StripLeft(constant.Example))
	assert.String(t, "example.org", host.StripLeft("test.example.org"))
	assert.String(
		t,
		"test.example.org",
		host.StripLeft("test2.test.example.org"),
	)
}

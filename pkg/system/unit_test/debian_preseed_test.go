package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/debian/preseed"
	"testing"
)

func TestDebianPreseedLink(t *testing.T) {
	assert.String(
		t,
		"https://www.debian.org/releases/buster/example-preseed.txt",
		preseed.Link("buster"),
	)
}

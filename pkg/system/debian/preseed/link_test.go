package preseed

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"https://www.debian.org/releases/buster/example-preseed.txt",
		Link("buster"),
	)
}

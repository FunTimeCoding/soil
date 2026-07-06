package join

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestJoin(t *testing.T) {
	assert.String(t, "a/b/c", Join("a", "b", "c"))
	assert.String(t, "/a/b/c", Join("/a", "b", "c"))
	assert.String(t, "a/b/c", Join("a", "/b", "c"))
	assert.String(t, "a/b/c", Join("a", "b", "/c"))
	assert.String(t, "/a/b/c", Join("/a", "/b", "c"))
	assert.String(t, "a/b/c", Join("a", "/b", "/c"))
	assert.String(t, "/a/b/c", Join("/a", "b", "/c"))
	assert.String(t, "/a/b/c", Join("/a/", "/b/", "/c/"))
}

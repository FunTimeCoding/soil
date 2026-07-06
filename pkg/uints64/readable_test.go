package uints64

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestReadable(t *testing.T) {
	assert.String(t, "1.0 KB", Readable(1024))
	assert.String(t, "1.0 MB", Readable(1048576))
	assert.String(t, "1.0 GB", Readable(1073741824))
	assert.String(t, "1.0 TB", Readable(1099511627776))
}

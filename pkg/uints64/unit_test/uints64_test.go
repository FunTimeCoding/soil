package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/uints64"
	"testing"
)

func TestReadable(t *testing.T) {
	assert.String(t, "1.0 KB", uints64.Readable(1024))
	assert.String(t, "1.0 MB", uints64.Readable(1048576))
	assert.String(t, "1.0 GB", uints64.Readable(1073741824))
	assert.String(t, "1.0 TB", uints64.Readable(1099511627776))
}

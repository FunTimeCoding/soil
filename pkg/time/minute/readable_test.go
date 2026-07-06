package minute

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestReadable(t *testing.T) {
	assert.String(t, "1 minute", Readable(1))
	assert.String(t, "1 hour", Readable(60))
}

package maps

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	assert.Bytes(t, []byte(`{"a":"b"}`), Encode(map[string]string{"a": "b"}))
}

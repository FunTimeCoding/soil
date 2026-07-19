package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/maps"
	"testing"
)

func TestEncode(t *testing.T) {
	assert.Bytes(t, []byte(`{"a":"b"}`), maps.Encode(map[string]string{"a": "b"}))
}

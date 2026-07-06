package message

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query_result"
	"testing"
)

func TestMessage(t *testing.T) {
	r, v := NewSlice(query_result.New())
	assert.NotNil(t, r)
	assert.NotNil(t, v)
}

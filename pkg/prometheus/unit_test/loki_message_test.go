package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query_result"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	"testing"
)

func TestMessage(t *testing.T) {
	r, v := message.NewSlice(query_result.New())
	assert.NotNil(t, r)
	assert.NotNil(t, v)
}

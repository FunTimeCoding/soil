package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/round_tripper"
	"testing"
)

func TestRoundTripper(t *testing.T) {
	assert.NotNil(t, round_tripper.New("", ""))
}

package metric

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestMetric(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa))
}

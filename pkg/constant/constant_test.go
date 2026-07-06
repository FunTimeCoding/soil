package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "go", Go)
	assert.String(t, "tag", TagKey)
	assert.String(t, "label", LabelKey)
	assert.String(t, "00:00:00:00:00:01", PhysicalTest1)
	assert.String(t, "00:00:00:00:00:02", PhysicalTest2)
	assert.String(t, "Unauthorized", Unauthorized)
}

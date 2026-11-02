package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "go", soil.Go)
	assert.String(t, "tag", soil.TagKey)
	assert.String(t, "label", soil.LabelKey)
	assert.String(t, "00:00:00:00:00:01", soil.PhysicalTest1)
	assert.String(t, "00:00:00:00:00:02", soil.PhysicalTest2)
	assert.String(t, "Unauthorized", soil.Unauthorized)
}

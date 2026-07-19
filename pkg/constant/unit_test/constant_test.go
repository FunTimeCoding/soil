package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "go", constant.Go)
	assert.String(t, "tag", constant.TagKey)
	assert.String(t, "label", constant.LabelKey)
	assert.String(t, "00:00:00:00:00:01", constant.PhysicalTest1)
	assert.String(t, "00:00:00:00:00:02", constant.PhysicalTest2)
	assert.String(t, "Unauthorized", constant.Unauthorized)
}

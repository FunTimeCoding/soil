package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"testing"
)

func TestItemConstant(t *testing.T) {
	assert.Count(t, 15, constant.Collectors)
}

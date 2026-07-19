package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Count(t, 3, constant.Severities)
	assert.Count(t, 4, constant.Statuses)
}

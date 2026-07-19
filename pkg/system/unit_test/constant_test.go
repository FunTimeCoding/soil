package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "src", constant.Source)
	assert.String(t, ".osquery/shell.em", constant.QuerySocketPath)
}

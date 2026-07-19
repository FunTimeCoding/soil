package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/pushover/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "PUSHOVER_TOKEN", constant.TokenEnvironment)
}

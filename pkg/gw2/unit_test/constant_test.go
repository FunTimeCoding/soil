package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "GW2_TOKEN", constant.TokenEnvironment)
}

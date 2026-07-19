package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/ssh/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "-T", constant.NoPTYArgument)
	assert.String(t, "-tt", constant.ForcePTYArgument)
	assert.String(t, "-v", constant.VerboseArgument)
}

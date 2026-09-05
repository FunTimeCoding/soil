package unit

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestValidateArgument(t *testing.T) {
	run.ValidateArgument(constant.UpperAlfa)
}

package unit

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestValidateArgument(t *testing.T) {
	run.ValidateArgument(stringsConstant.UpperAlfa)
}

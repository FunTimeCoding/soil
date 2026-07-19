package unit_test

import (
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestValidateArgument(t *testing.T) {
	run.ValidateArgument(upper.Alfa)
}

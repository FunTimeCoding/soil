package unit_test

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestValidateArchitecture(t *testing.T) {
	system.ValidateArchitecture(constant.AMD64)
}

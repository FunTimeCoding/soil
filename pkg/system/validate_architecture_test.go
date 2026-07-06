package system

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestValidateArchitecture(t *testing.T) {
	ValidateArchitecture(constant.AMD64)
}

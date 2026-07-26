package unit_test

import (
	"github.com/funtimecoding/soil/pkg/kubernetes"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestValidateName(t *testing.T) {
	kubernetes.ValidateName(constant.LowerAlfa)
}

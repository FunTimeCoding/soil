package kubernetes

import (
	"github.com/funtimecoding/soil/pkg/strings/lower"
	"testing"
)

func TestValidateName(t *testing.T) {
	ValidateName(lower.Alfa)
}

package git

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestStatus(t *testing.T) {
	Status(system.ParentDirectory(constant.Depth))
}

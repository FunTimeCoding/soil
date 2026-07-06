package git

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestTree(t *testing.T) {
	Tree(system.ParentDirectory(constant.Depth))
}

package git

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestTags(t *testing.T) {
	Tags(system.ParentDirectory(constant.Depth))
}

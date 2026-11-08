package face_tester

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"go/types"
	"golang.org/x/tools/go/packages"
	"testing"
)

type Tester struct {
	Loaded    []*packages.Package
	Workspace *index.Workspace
	model     *types.Named
	t         *testing.T
}

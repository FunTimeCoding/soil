package face_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/source/unit/kind_tester"
	"go/types"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	directory := kind_tester.ImporterModule(t)
	result, _, e := resolve.LoadPackages(directory, "./...")
	errors.PanicOnError(e)
	loaded := resolve.PreferTestVariants(result)
	var model *types.Named

	for _, p := range loaded {
		if o := p.Types.Scope().Lookup("Model"); o != nil {
			model = o.Type().(*types.Named)
		}
	}

	if model == nil {
		t.Fatalf("fixture module declares no Model")
	}

	return &Tester{
		Loaded:    loaded,
		Workspace: index.New(t.TempDir(), directory, face.Kind()),
		model:     model,
		t:         t,
	}
}

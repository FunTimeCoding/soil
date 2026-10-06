package reference_tester

import (
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/types"
	"testing"
)

func Target(
	t *testing.T,
	directory string,
	typeName string,
	member string,
) string {
	t.Helper()
	loaded, e := resolve.LoadBase(directory, "./alfa")

	if e != nil || len(loaded) == 0 {
		t.Fatalf("load alfa: %v", e)
	}

	o := loaded[0].Types.Scope().Lookup(typeName)

	if member != "" {
		o, _, _ = types.LookupFieldOrMethod(
			types.NewPointer(o.Type()),
			true,
			loaded[0].Types,
			member,
		)
	}

	result, okay := xref.Target(o)

	if !okay {
		t.Fatalf("no target for %s.%s", typeName, member)
	}

	return result
}

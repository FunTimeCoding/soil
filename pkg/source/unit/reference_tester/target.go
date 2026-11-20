package reference_tester

import (
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/types"
)

func (o *Tester) Target(
	directory string,
	typeName string,
	member string,
) string {
	o.t.Helper()
	loaded, e := resolve.LoadBase(directory, "./alfa")

	if e != nil || len(loaded) == 0 {
		o.t.Fatalf("load alfa: %v", e)
	}

	object := loaded[0].Types.Scope().Lookup(typeName)

	if member != "" {
		object, _, _ = types.LookupFieldOrMethod(
			types.NewPointer(object.Type()),
			true,
			loaded[0].Types,
			member,
		)
	}

	result, okay := xref.Target(object)

	if !okay {
		o.t.Fatalf("no target for %s.%s", typeName, member)
	}

	return result
}

package xref

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
	"golang.org/x/tools/go/types/objectpath"
)

func Target(o types.Object) (string, bool) {
	switch v := o.(type) {
	case *types.Func:
		o = v.Origin()
	case *types.Var:
		o = v.Origin()
	}

	if o.Pkg() == nil {
		return "", false
	}

	path, e := objectpath.For(o)

	if e != nil {
		return "", false
	}

	return join.Space(o.Pkg().Path(), string(path)), true
}

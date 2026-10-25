package nil_dependency

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
)

func dependencyFace(t types.Type) (string, bool) {
	named, okay := t.(*types.Named)

	if !okay || named.Obj().Pkg() == nil {
		return "", false
	}

	if _, okay = named.Underlying().(*types.Interface); !okay {
		return "", false
	}

	if !isDependencyPath(named.Obj().Pkg().Path()) {
		return "", false
	}

	return join.Dot(
		[]string{named.Obj().Pkg().Name(), named.Obj().Name()},
	), true
}

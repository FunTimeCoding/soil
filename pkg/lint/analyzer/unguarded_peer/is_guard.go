package unguarded_peer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func isGuard(t types.Type) bool {
	if pointer, okay := t.(*types.Pointer); okay {
		t = pointer.Elem()
	}

	named, okay := t.(*types.Named)

	return okay &&
		named.Obj().Pkg() != nil &&
		named.Obj().Pkg().Path() == constant.PeerPackage &&
		named.Obj().Name() == constant.GuardType
}

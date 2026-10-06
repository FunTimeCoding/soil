package service

import "go/types"

func sameNamed(
	t types.Type,
	named *types.Named,
) bool {
	if types.Identical(t, named) {
		return true
	}

	other, okay := t.(*types.Named)

	if !okay {
		return false
	}

	return other.Obj().Pos().IsValid() &&
		other.Obj().Pos() == named.Obj().Pos() &&
		other.Obj().Name() == named.Obj().Name()
}

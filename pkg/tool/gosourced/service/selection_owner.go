package service

import "go/types"

func selectionOwner(selection *types.Selection) string {
	subject := selection.Recv()

	if pointer, okay := subject.(*types.Pointer); okay {
		subject = pointer.Elem()
	}

	if named, okay := subject.(*types.Named); okay {
		return named.Obj().Name()
	}

	return ""
}

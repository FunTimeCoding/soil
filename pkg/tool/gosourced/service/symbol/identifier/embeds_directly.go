package identifier

import "go/types"

func embedsDirectly(
	t types.Type,
	typeName *types.TypeName,
) bool {
	if p, okay := t.(*types.Pointer); okay {
		t = p.Elem()
	}

	n, okay := t.(*types.Named)

	return okay && n.Origin().Obj().Pos() == typeName.Pos()
}
